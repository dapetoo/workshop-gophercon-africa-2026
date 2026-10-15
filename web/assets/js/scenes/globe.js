import { mount } from '/assets/js/threejs.js';

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

mount('globe', function ({ THREE, scene, camera, colour, pointer, canvas }) {
	const globe = new THREE.Group();
	scene.add(globe);

	const radius = 1.6;

	const earth = new THREE.Mesh(
		new THREE.SphereGeometry(radius, 64, 64),
		new THREE.MeshStandardMaterial({ color: colour('--three-primary', '#2f5a41'), roughness: 0.9, metalness: 0 }),
	);
	globe.add(earth);

	// Lines of latitude and longitude, like a field globe.
	const grid = new THREE.LineSegments(
		new THREE.WireframeGeometry(new THREE.SphereGeometry(radius * 1.002, 24, 16)),
		new THREE.LineBasicMaterial({ color: colour('--three-grid', '#4d7a5c'), transparent: true, opacity: 0.35 }),
	);
	globe.add(grid);

	const positions = [];
	for (const pair of readPoints()) {
		const lat = Number(pair[0]);
		const lng = Number(pair[1]);
		if (!Number.isFinite(lat) || !Number.isFinite(lng)) continue;
		const v = toVector(THREE, lat, lng, radius * 1.01);
		positions.push(v.x, v.y, v.z);
	}
	const geometry = new THREE.BufferGeometry();
	geometry.setAttribute('position', new THREE.Float32BufferAttribute(positions, 3));
	const dots = new THREE.Points(
		geometry,
		new THREE.PointsMaterial({ color: colour('--three-accent', '#b8e04a'), size: 0.06, sizeAttenuation: true }),
	);
	globe.add(dots);

	// A soft halo around the rim.
	const halo = new THREE.Mesh(
		new THREE.SphereGeometry(radius * 1.12, 48, 48),
		new THREE.MeshBasicMaterial({ color: colour('--three-accent', '#b8e04a'), transparent: true, opacity: 0.06, side: THREE.BackSide }),
	);
	scene.add(halo);

	scene.add(new THREE.AmbientLight(0xffffff, 0.9));
	const sun = new THREE.DirectionalLight(0xfff4d6, 2.2);
	sun.position.set(-3, 3, 5);
	scene.add(sun);
	const rim = new THREE.DirectionalLight(colour('--three-accent', '#b8e04a'), 1.5);
	rim.position.set(4, -1, -3);
	scene.add(rim);

	// Sit the globe to the right on wide screens, centred behind the text on phones.
	function place() {
		const wide = canvas.clientWidth > 860;
		globe.position.x = wide ? 1.9 : 0;
		globe.position.y = wide ? 0 : -1.4;
		halo.position.copy(globe.position);
	}
	place();
	window.addEventListener('resize', place, { passive: true });

	globe.rotation.x = 0.35;
	camera.position.z = 5.2;

	return function (time) {
		globe.rotation.y = time * 0.08 + pointer.x * 0.25;
		globe.rotation.x = 0.35 + pointer.y * 0.1;
	};
});
