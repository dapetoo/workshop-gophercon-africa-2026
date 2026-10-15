import { mount } from '/assets/js/threejs.js';

// three.js's own land-and-water map: water is white, land is black.
const LAND_MAP = 'https://cdn.jsdelivr.net/gh/mrdoob/three.js@r170/examples/textures/planets/earth_specular_2048.jpg';

// Sightings as [lat, lng] pairs, written into the page by the home template.
function readPoints() {
	const el = document.getElementById('globe-points');
	if (!el) return [];
	try {
		const data = JSON.parse(el.textContent);
		return Array.isArray(data) ? data : [];
	} catch {
		return [];
	}
}

function toVector(THREE, lat, lng, r) {
	const la = (lat * Math.PI) / 180;
	const lo = (lng * Math.PI) / 180;
	return new THREE.Vector3(r * Math.cos(la) * Math.sin(lo), r * Math.sin(la), r * Math.cos(la) * Math.cos(lo));
}

// Repaints the land map in the site's colours, so the globe matches the page.
function paintEarth(THREE, image, land, water) {
	const canvas = document.createElement('canvas');
	canvas.width = 1024;
	canvas.height = 512;
	const ctx = canvas.getContext('2d', { willReadFrequently: true });
	ctx.drawImage(image, 0, 0, canvas.width, canvas.height);
	const px = ctx.getImageData(0, 0, canvas.width, canvas.height);
	const d = px.data;
	const l = {}, w = {};
	land.getRGB(l, THREE.SRGBColorSpace);
	water.getRGB(w, THREE.SRGBColorSpace);
	for (let i = 0; i < d.length; i += 4) {
		const t = d[i] / 255;
		d[i] = (l.r * (1 - t) + w.r * t) * 255;
		d[i + 1] = (l.g * (1 - t) + w.g * t) * 255;
		d[i + 2] = (l.b * (1 - t) + w.b * t) * 255;
	}
	ctx.putImageData(px, 0, 0);
	const texture = new THREE.CanvasTexture(canvas);
	texture.colorSpace = THREE.SRGBColorSpace;
	return texture;
}

mount('globe', function ({ THREE, scene, camera, colour, pointer, still }) {
	const globe = new THREE.Group();
	scene.add(globe);

	const radius = 1.7;
	const water = colour('--three-primary', '#2f5a41');
	const land = colour('--three-land', '#e9e3cf');

	const material = new THREE.MeshStandardMaterial({ color: water, roughness: 0.85, metalness: 0 });
	const earth = new THREE.Mesh(new THREE.SphereGeometry(radius, 96, 96), material);
	// The texture's longitude 0 sits on +x; the dots put it on +z.
	earth.rotation.y = -Math.PI / 2;
	globe.add(earth);

	const image = new Image();
	image.crossOrigin = 'anonymous';
	image.onload = function () {
		material.map = paintEarth(THREE, image, land, water);
		material.color.set(0xffffff);
		material.needsUpdate = true;
	};
	image.src = LAND_MAP;

	// Faint lines of latitude and longitude, like a field globe.
	const grid = new THREE.LineSegments(
		new THREE.WireframeGeometry(new THREE.SphereGeometry(radius * 1.003, 24, 12)),
		new THREE.LineBasicMaterial({ color: land, transparent: true, opacity: 0.08 }),
	);
	globe.add(grid);

	// A small round dot texture, so points aren't squares.
	const dot = document.createElement('canvas');
	dot.width = dot.height = 64;
	const g = dot.getContext('2d');
	const grad = g.createRadialGradient(32, 32, 0, 32, 32, 32);
	grad.addColorStop(0, 'rgba(255,255,255,1)');
	grad.addColorStop(0.45, 'rgba(255,255,255,1)');
	grad.addColorStop(0.55, 'rgba(255,255,255,0.35)');
	grad.addColorStop(1, 'rgba(255,255,255,0)');
	g.fillStyle = grad;
	g.fillRect(0, 0, 64, 64);

	const positions = [];
	for (const pair of readPoints()) {
		const lat = Number(pair[0]);
		const lng = Number(pair[1]);
		if (!Number.isFinite(lat) || !Number.isFinite(lng)) continue;
		const v = toVector(THREE, lat, lng, radius * 1.012);
		positions.push(v.x, v.y, v.z);
	}
	const geometry = new THREE.BufferGeometry();
	geometry.setAttribute('position', new THREE.Float32BufferAttribute(positions, 3));
	const dots = new THREE.Points(
		geometry,
		new THREE.PointsMaterial({
			color: colour('--three-accent', '#b8e04a'),
			size: 0.11,
			map: new THREE.CanvasTexture(dot),
			transparent: true,
			depthWrite: false,
		}),
	);
	globe.add(dots);

	// A soft lime halo around the rim.
	const halo = new THREE.Mesh(
		new THREE.SphereGeometry(radius * 1.14, 64, 64),
		new THREE.ShaderMaterial({
			uniforms: { glow: { value: colour('--three-accent', '#b8e04a') } },
			vertexShader: 'varying vec3 vN; void main(){ vN = normalize(normalMatrix * normal); gl_Position = projectionMatrix * modelViewMatrix * vec4(position,1.0); }',
			fragmentShader: 'uniform vec3 glow; varying vec3 vN; void main(){ float i = pow(0.62 - dot(vN, vec3(0.0,0.0,1.0)), 3.0); gl_FragColor = vec4(glow, 1.0) * i; }',
			side: THREE.BackSide,
			blending: THREE.AdditiveBlending,
			transparent: true,
			depthWrite: false,
		}),
	);
	scene.add(halo);

	scene.add(new THREE.AmbientLight(0xffffff, 2.6));
	const sun = new THREE.DirectionalLight(0xfff4d6, 1.1);
	sun.position.set(-4, 3, 5);
	scene.add(sun);

	// Start over Africa, tilted like a desk globe.
	const start = -0.35;
	globe.rotation.set(0.3, start, 0);
	camera.position.z = 5.4;

	return function (time) {
		const spin = still ? 0 : time * 0.06;
		globe.rotation.y = start - spin + pointer.x * 0.2;
		globe.rotation.x = 0.3 + pointer.y * 0.08;
	};
});
