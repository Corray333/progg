<template>
    <div class="card" ref="card" v-if="props.pick!=undefined" @click="console.log('test')">
        <span class="card-header">
            <img :src="require(`../assets/${props.pick.img}`)" alt="">
            <h3>{{ props.pick.name }}</h3>
        </span>
        <div v-if="props.pick.color != undefined" :class="props.pick.color"></div>
        <p>Владелец: {{ props.pick.owner }}</p>
        <p>Цена: {{ props.pick.price }}</p>
        <p>Рента: {{ props.pick.rentLevels[props.pick.progsCount] }}</p>
    </div>
</template>

<script setup>
import { defineProps, ref } from 'vue'
import {useEventListener} from '@vueuse/core'
const props = defineProps(['pick'])

const card = ref(null)

useEventListener(card, 'mousemove', (e)=>{
    var rect = card.value.getBoundingClientRect();
    let xAxis = ((rect.left+112) - e.pageX)/7
    let yAxis = ((rect.top+150) - e.pageY)/7
    card.value.style.transform = `scale(1.03) rotateY(${yAxis}deg) rotateX(${xAxis}deg) translateZ(50px)`
    card.value.style.transition = '0'
})

useEventListener(card, 'mouseleave', ()=>{
    card.value.style.transform = `rotateY(0deg) rotateX(0deg)`
    card.value.style.transition = '0.3s'
})



</script>


<style scoped>

.card{
    width: 225px;
    height: 300px;
    padding: 5px;
    border: 3px solid white;
    border-radius: 15px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 5px;
    transition: ease-in-out 0.3s;
    background-color: var(--black);
}
.card:hover{
    transform: scale(1.1);
    transition: ease-in-out 0.3s;
    box-shadow: 0 0 15px 0 rgba(255,255,255,0.5);
}
.card-header{
    display: flex;
    align-items: center;
    border-radius: 15px;
    padding: 5px;
}
.card-header>img{
    height: 32px;
}
.card>div {
    width: 40px;
    height: 5px;
    border-radius: 99px;
    margin: 0 20px;
}
.blue{
    background-color: var(--blue);
}
.yellow{
    background-color: var(--yellow);
}
.white{
    background-color: white;
}
.red{
    background-color: var(--red);
}
.green{
    background-color: var(--green);
}
.purple{
    background-color: var(--purple);
}
.orange{
    background-color: var(--orange);
}
.light-blue{
    background-color: var(--light-blue);
}

</style>