<template>
    <div class="card" ref="card" v-if="pick!=undefined" @click="console.log('test')" key="card">
        <span class="card-header">
            <img :src="`/src/assets/icons/${pick.key}.png`" alt="">
            <h3>{{ pickRef.name }}</h3>
        </span>
        <div class="line" v-if="pick.color != undefined" :class="pick.color"></div>
        <p>Владелец: {{ pickRef.owner }}</p>
        <p>Цена: {{ pick.price }}</p>
        <p>Рента: {{ pick.rentLevels[pick.progsCount] }}</p>
        <div class="row">
            <img src="/src/assets/icons/prog.png" alt="" v-for="i in pick.progsCount" :key="i">
        </div>
        <h3>Игроки на поле</h3>
        <div class="players">
            <div v-for="(player,i) of players"  :key="i" >
                <img  class="player" v-if="player.position == pick.number" :src="`/src/assets/avatars/${player.avatar}.png`" alt="">
            </div>
        </div>
        <button @click="$emit('buy', pick.key)" v-if="pick.owner == 'отсутствует' && activePlayer != undefined && playerProfile != undefined && activePlayer.username == playerProfile.username && playerProfile.position == pick.number">Купить</button>
    </div>
</template>

<script setup>
import { ref, toRef } from 'vue'
import {useEventListener} from '@vueuse/core'
const props = defineProps(['pick', 'players', 'activePlayer', 'playerProfile'])

const card = ref(null)

const pickRef = toRef(props, 'pick')


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
.line {
    width: 40px;
    height: 5px;
    border-radius: 99px;
    margin: 0 20px;
}
.row{
    display: flex;
    gap: 5px;
}
.row>img{
    height: 24px;
    padding: 10px;
    border: 3px solid white;
    border-radius: 99px;    
    box-sizing: content-box;
}

.players{
    display: flex;
    gap: 5px;
}
.player{
    width: 50px;
    border:3px solid white;
    border-radius: 99px;
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