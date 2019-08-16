
var container = document.getElementById('num1');

var renderer = new THREE.WebGLRenderer(),
    containerWidth1 = container.offsetWidth,
    containerHeight1 = container.offsetHeight;

renderer.setSize(containerWidth1, containerHeight1);

container.appendChild(renderer.domElement);

var geometry = new THREE.BoxGeometry(1, 1, 1);
var material = new THREE.MeshBasicMaterial({ color: 0x00ff00 });
var cube = new THREE.Mesh(geometry, material);

var scene = new THREE.Scene();
scene.add(cube);

var camera = new THREE.PerspectiveCamera(75, container.offsetWidth / container.offsetHeight, 0.1, 1000);
camera.position.z = 5;

//second

var container2 = document.getElementById('num2'),
    range = 50,
    containerWidth = container2.offsetWidth,
    containerHeight = container2.offsetHeight;

var renderer2 = new THREE.WebGLRenderer();
renderer2.setSize(containerWidth, containerHeight);
container2.appendChild(renderer2.domElement);

var camera2 = new THREE.PerspectiveCamera(75, containerWidth / containerHeight, 0.1, 1000);
//camera2.position.z = 55;
camera2.position.set(0, 0, range);
camera2.lookAt(new THREE.Vector3(0, 0, 0));

var scene2 = new THREE.Scene();


var geometry2 = new THREE.SphereGeometry(1, 8, 6);
spheres = new THREE.Object3D();
scene2.add(spheres);

for (var i = 0; i < 100; i++) {
    var grayness = Math.random() * 0.5 + 0.25,
        mat = new THREE.MeshBasicMaterial(),
        sphere = new THREE.Mesh(geometry2, mat);
    mat.color.setRGB(grayness, grayness, grayness);
    sphere.position.set(range * (0.5 - Math.random()), range * (0.5 - Math.random()), range * (0.5 - Math.random()));
    //sphere.rotation.set(Math.random(), Math.random(), Math.random()).multiplyScalar(2 * Math.PI);
    sphere.grayness = grayness;
    spheres.add(sphere);
}

var raycaster = new THREE.Raycaster(),
    mouseVector = new THREE.Vector2(),
    obj,
    content = document.getElementById('content');

container2.addEventListener('mousemove', onMouseMove, false);
container2.addEventListener('click', onClick, false);

function onMouseMove(e) {

    obj = null;

    mouseVector.x = 2 * (e.offsetX / containerWidth) - 1;
    mouseVector.y = 1 - 2 * (e.offsetY / containerHeight);

    raycaster.setFromCamera(mouseVector, camera2)

    var intersects = raycaster.intersectObjects(spheres.children);


    spheres.children.forEach(function (sphere) {
        sphere.material.color.setRGB(sphere.grayness, sphere.grayness, sphere.grayness);
    });


    for (var i = 0; i < intersects.length; i++) {
        obj = intersects[0].object;
        console.log(intersects);
        obj.material.color.set(0xff0000);
    }


}

function onClick(e) {
    if (obj != null) {
        content.style.display = "block"
        content.style.top = e.clientY;
        content.style.left = e.clientX;
    }
    else content.style.display = "none"
}



//render

var animate = function () {
    requestAnimationFrame(animate);

    cube.rotation.x += 0.01;
    cube.rotation.y += 0.01;

    renderer.render(scene, camera);
    renderer2.render(scene2, camera2);
};

animate();
