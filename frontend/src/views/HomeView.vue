<template>
  <div class="home">
    <canvas ref="canvasRef"></canvas>
  </div>
</template>

<script setup>
import * as THREE from "three"
import { ref, onMounted } from "vue"
import {OrbitControls} from "three/examples/jsm/controls/OrbitControls.js"
import {GLTFLoader} from "three/examples/jsm/loaders/GLTFLoader.js"
import {DRACOLoader} from "three/examples/jsm/loaders/DRACOLoader.js"
import { RGBELoader } from 'three/addons/loaders/RGBELoader.js';
import {RenderPass} from "three/examples/jsm/postprocessing/RenderPass.js"
import {EffectComposer} from "three/examples/jsm/postprocessing/EffectComposer.js"
import {UnrealBloomPass} from "three/examples/jsm/postprocessing/UnrealBloomPass.js"
import { SSAARenderPass } from 'three/examples/jsm/postprocessing/SSAARenderPass.js';


let scene = new THREE.Scene()
let renderer
let controls
let canvasRef = ref()

let shownModel
let loadedModel
let modelName = "City"



let camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 100)
camera.position.x=-1.5
camera.position.y=1.5
camera.position.z=2
scene.add(camera)


let renderScene
let composer

let ssaaRenderPass





let loop = () => {
  requestAnimationFrame(loop)
  composer.render()
}

let resizeCallback = () => {
  camera.aspect = window.innerWidth / window.innerHeight
  camera.updateProjectionMatrix()
  renderer.setSize(window.innerWidth, window.innerHeight)
}

onMounted(() => {
  renderer = new THREE.WebGLRenderer({
    canvas: canvasRef.value,
    alpha: true,
    antialias: true,
  })
  renderer.setSize(window.innerWidth, window.innerHeight)
  renderer.setPixelRatio(window.devicePixelRatio)
  renderer.render(scene, camera)


  const pmremGenerator = new THREE.PMREMGenerator( renderer )

  const hdriLoader = new RGBELoader()
  hdriLoader.load( '/texture3.hdr', function ( texture ) {
    const envMap = pmremGenerator.fromEquirectangular( texture ).texture
    texture.dispose()
    scene.environment = envMap
  })

  renderScene = new RenderPass( scene, camera );
  composer = new EffectComposer( renderer );
  composer.setPixelRatio( 1 );
  composer.addPass( renderScene );


  ssaaRenderPass = new SSAARenderPass( scene, camera );
  ssaaRenderPass.sampleLevel = 4;
  composer.addPass( ssaaRenderPass );

  const bloomPass = new UnrealBloomPass(
    new THREE.Vector2( window.innerWidth, window.innerHeight ),
    0.3,
    1,
    0
  )
  composer.addPass( bloomPass)


  
  loop()
  controls = new OrbitControls(camera, canvasRef.value)
  controls.enableDamping = true








  const gltfLoader = new GLTFLoader()
  const dracoLoader = new DRACOLoader()
  dracoLoader.setDecoderPath("/draco3d/")
  gltfLoader.setDRACOLoader(dracoLoader)
  gltfLoader.load("/models/progg.gltf", (gltf) => {
    loadedModel = gltf.scene
    shownModel = loadedModel.clone()
    shownModel.name = modelName
    shownModel.scale.set(0.15,0.15,0.15)
    scene.add(shownModel)

    
  })

  window.addEventListener("resize", resizeCallback)
})

</script>
<style>
.home{
  height: 100vh;
  background-color: #1b1b1b;
}
canvas{
  width: 500px;
  height: 500px;
}
</style>
