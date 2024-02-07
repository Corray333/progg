<template>
    <div class="room">
        <transition>
            <div class="quiz-wrapper" v-if="quiz!=null">
                <QuizModal :quiz="quiz" :active-player="activePlayer" :player-profile="playerProfile" @answer="answer"/>
            </div>
        </transition>

        <div class="join-room-modal-wrapper" v-show="showModal">
            <div class="join-room-modal">
                <h2>Войти в комнату</h2>
                <span>
                    <p>Имя пользователя: </p>
                    <input type="text" v-model="username" autocomplite="off">
                </span>
                <span>
                    <p>Пароль: </p>
                    <input type="password" v-model="password" autocomplite="off">
                </span>
                <button @click="joinRoom()">Войти</button>
                <button id="close-modal" @click="showModal = false; router.push('/rooms')"><img src="../assets/icons/plus.png" alt=""></button>
            </div>
        </div>
        <PlayMap @tile-pick="changeCard" @move="move" :players="players" :endOfTurn="endOfTurn" :playerProfile="playerProfile" ref="mapComponent"/>
        <div class="column">
            <div class="row">
                <InfoCard :key="pick != undefined ? pick.owner : 'card'" @buy="buy" :pick="pick" :players="players" :activePlayer="activePlayer" :playerProfile="playerProfile"/>
                <PlayersList :active-player="activePlayer" @stopTheTurn="stopTheTurn" :players="players" :playerProfile="playerProfile"  @ready="ready" :turnTimer="turnTimer"/>
            </div>
            <HandCards/>
        </div>
    </div>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router'
import {ref, watch} from 'vue'
import axios from 'axios'


import PlayMap from '../components/PlayMap.vue'
import InfoCard from '../components/InfoCard.vue'
import PlayersList from '../components/PlayersList.vue'
import HandCards from '../components/HandCards.vue'
import QuizModal from '../components/QuizModal.vue'

const route = useRoute().params.room
const router = useRouter()

const showModal = ref(true)
const username = ref('')
const password = ref('')
let socket

const endOfTurn = ref()
const turnTimer = ref()
let timerInterval

let playerProfile = ref()
let activePlayer = ref()

watch(endOfTurn, ()=>{
    timer()
})

const timer = ()=>{
    turnTimer.value = Math.floor((endOfTurn.value-Date.now())/1000)
    clearInterval(timerInterval)
    timerInterval = setInterval(()=>{
        turnTimer.value--
    }, 1000)
}

const ready = ()=>{
    socket.send('01')
    playerProfile.ready = true
}

// TODO: fix yandex and sber
// Map with info about map
const map = new Map()
{
map.set('vk', {
        name:'ВКонтакте',
        description:'ВКонтакте — российская социальная сеть, предназначенная для общения между пользователями сети Интернет. Создана Дуровым Павлом Валерьевичем в 2006 году. В 2014 году владельцем социальной сети стала компания «ВКонтакте».',
        price:400,
        color:'blue',
        programs:['ВКонтакте','Одноклассники','ВК знакомства'],
        progsCount: 0,
        rentLevels: [50, 200, 600, 1400],
        progPrice: 200,
        owner: 'отсутствует',
        playersHere: [],
        key: 'vk',
        number: 36
    }
)
map.set('ozon', {
        name:'Ozon',
        description:'Ozon — российская интернет-компания, владеющая одноимённым интернет-магазином. Штаб-квартира компании расположена в Москве. Основана в 1998 году Дмитрием Бакшеевым и Александром Шульгином.',
        price:350,
        color:'blue',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [35, 175, 500, 1100],
        progPrice: 200,
        owner: 'отсутствует',
        playersHere: [],
        key: 'ozon',
        number: 34
    }
)
map.set('yandex', {
        name:'Yandex',
        description:'Yandex — российская транснациональная компания, владеющая одноимённой системой поиска в Сети, интернет-порталами и службами в нескольких странах. Штаб-квартира — в Москве.',
        price:300,
        color:'red',
        programs:['Yandex','Yandex музыка','Yandex облако'],
        progsCount: 0,
        rentLevels: [28, 250, 450, 1000],
        progPrice: 200,
        owner: 'отсутствует',
        playersHere: [],
        key: 'yandex',
        number: 30
    }
)
map.set('alpha', {
        name:'Alpha',
        description:'Ozon — российская интернет-компания, владеющая одноимённым интернет-магазином. Штаб-квартира компании расположена в Москве. Основана в 1998 году Дмитрием Бакшеевым и Александром Шульгином.',
        price:320,
        color:'red',
        programs:['Aplha банк','Alpha страхование','Alpha инвестиции'],
        progsCount: 0,
        rentLevels: [26, 130, 390, 900],
        progPrice: 200,
        owner: 'отсутствует',
        playersHere: [],
        key: 'alpha',
        number: 32
    }
)
map.set('magnit', {
        name:'Магнит',
        description:'',
        price:300,
        color:'red',
        programs:['Онлайн каталог','Доставка','Приложение для персонала'],
        progsCount: 0,
        rentLevels: [26, 130, 390, 900],
        progPrice: 150,
        owner: 'отсутствует',
        playersHere: [],
        key: 'magnit',
        number: 29
    }
)
map.set('rostelekom', {
        name:'Ростелеком',
        description:'',
        price:280,
        color:'purple',
        programs:['Wink','Ростелеком лицей','Умный дом'],
        progsCount: 0,
        rentLevels: [24, 120, 360, 850],
        progPrice: 150,
        owner: 'отсутствует',
        playersHere: [],
        key: 'rostelekom',
        number: 25
    }
)
map.set('netbynet', {
        name:'NetByNet',
        description:'',
        price:260,
        color:'purple',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [22, 110, 330, 800],
        progPrice: 150,
        owner: 'отсутствует',
        playersHere: [],
        key: 'netbynet',
        number: 27
    }
)
map.set('tinkoff', {
        name:'Tinkoff',
        description:'',
        price:240,
        color:'yellow',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [20, 100, 300, 750],
        progPrice: 150,
        owner: 'отсутствует',
        playersHere: [],
        key: 'tinkoff',
        number: 23
    }
)
map.set('beeline', {
        name:'Beeline',
        description:'',
        price:220,
        color:'yellow',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [20, 100, 300, 750],
        progPrice: 150,
        owner: 'отсутствует',
        playersHere: [],
        key: 'beeline',
        number: 22
    }
)
map.set('1c', {
        name:'1С',
        description:'',
        price:220,
        color:'yellow',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [18, 90, 250, 700],
        progPrice: 150,
        owner: 'отсутствует',
        playersHere: [],
        key: '1c',
        number: 20
    }
)
map.set('rostech', {
        name:'Ростех',
        description:'',
        price:200,
        color:'white',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [16, 80, 220, 600],
        progPrice: 100,
        owner: 'отсутствует',
        playersHere: [],
        key: 'rostech',
        number: 18
    }
)
map.set('astra', {
        name:'Астра',
        description:'',
        price:180,
        color:'white',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [14, 70, 200, 550],
        progPrice: 100,
        owner: 'отсутствует',
        playersHere: [],
        key: 'astra',
        number: 16
    }
)
map.set('sber', {
        name:'Сбер',
        description:'',
        price:160,
        color:'green',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [12, 60, 180, 500, 700],
        progPrice: 100,
        owner: 'отсутствует',
        playersHere: [],
        key: 'sber',
        number: 14
    }
)
map.set('kaspersky', {
        name:'Касперский',
        description:'',
        price:140,
        color:'green',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [10, 50, 150, 450, 625],
        progPrice: 100,
        owner: 'отсутствует',
        playersHere: [],
        key: 'kaspersky',
        number: 13
    }
)
map.set('megafon', {
        name:'Мегафон',
        description:'',
        price:140,
        color:'green',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [10, 50, 150, 450, 625],
        progPrice: 100,
        owner: 'отсутствует',
        playersHere: [],
        key: 'megafon',
        number: 11
    }
)
map.set('finam', {
        name:'Финам',
        description:'',
        price:120,
        color:'orange',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [8, 40, 100, 300, 450],
        progPrice: 50,
        owner: 'отсутствует',
        playersHere: [],
        key: 'finam',
        number: 9
    }
)
map.set('tensor', {
        name:'Тензор',
        description:'',
        price:100,
        color:'orange',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [6, 30, 90, 270, 400],
        progPrice: 50,
        owner: 'отсутствует',
        playersHere: [],
        key: 'tensor',
        number: 8
    }
)
map.set('europlan', {
        name:'Европлан',
        description:'',
        price:100,
        color:'orange',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [6, 30, 90, 270, 400],
        progPrice: 50,
        owner: 'отсутствует',
        playersHere: [],
        key: 'europlan',
        number: 6
    }
)
map.set('yota', {
        name:'Yota',
        description:'',
        price:60,
        color:'light-blue',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [4, 20, 60, 180, 320],
        progPrice: 50,
        owner: 'отсутствует',
        playersHere: [],
        key: 'yota',
        number: 4
    }
)
map.set('skyeng', {
        name:'Sky-eng',
        description:'',
        price:60,
        color:'light-blue',
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [2, 10, 30, 90, 160],
        progPrice: 50,
        owner: 'отсутствует',
        playersHere: [],
        key: 'skyeng',
        number: 2
    }
)
map.set('hh', {
        name:'HeadHunter',
        description:'',
        price:200,
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [20, 50, 100, 200, 300],
        owner: 'отсутствует',
        playersHere: [],
        key: 'hh',
        number: 5
    }
)
map.set('geekbrains', {
        name:'Geekbrains',
        description:'',
        price:200,
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [20, 50, 100, 200, 300],
        owner: 'отсутствует',
        playersHere: [],
        key: 'geekbrains',
        number: 17
    }
)
map.set('wot', {
        name:'World of Tanks',
        description:'',
        price:200,
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [20, 50, 100, 200, 300],
        owner: 'отсутствует',
        playersHere: [],
        key: 'wot',
        number: 26
    }
)
map.set('intellij', {
        name:'IntelliJ IDEA',
        description:'',
        price:200,
        programs:['Ozon','Ozon банк','OZON Travel'],
        progsCount: 0,
        rentLevels: [20, 50, 100, 200, 300],
        owner: 'отсутствует',
        playersHere: [],
        key: 'intellij',
        number: 35
    }
)
}



const pick = ref()
const changeCard = (query)=>{
    pick.value = map.get(query)
}

const quiz = ref(null)

const buy = (company)=>{
    if (map.get(company).price > playerProfile.money){
        alert('У вас недостаточно денег')
        return
    }
    socket.send('06'+(map.get(company).progPrice==undefined?'true':'false'))
}
const answer = (variant)=>{
    let req = {
        company: pick.value.key,
        answer: variant,
        quiz_id: quiz.value.quiz_id,
        is_game: pick.value.progPrice==undefined
    }
    socket.send('04'+JSON.stringify(req))
}

const players = ref()
const mapComponent = ref(null)

watch(players, ()=>{
    mapComponent.value.placePlayers()
}, {deep: true})

const stopTheTurn = ()=>{
    socket.send('07')
}

const joinRoom = async ()=>{
    try { 
        let req = {
            player_name:username.value,
            password:password.value
        }
        await axios.post(`http://localhost:3000/rooms/${route}`, req) 
        socket = new WebSocket(`ws://localhost:3000/rooms/${route}/game`);

        socket.onopen = function() {
            console.log("[open] Connection established");
            console.log("Sending to server");
            socket.send(username.value+'|'+route);
            socket.send('05'+username.value)
            socket.send('02')
            showModal.value = false
        };

        socket.onmessage = function(event) {
            handleFunc(event.data)
        };

        socket.onclose = function(event) {

            if (event.wasClean) {
                console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
            } else {
                alert('Другой пользователь вошел под тем же именем')
                console.log('[close] Connection died');
            }
            router.push(`/home`)
        }

        socket.onerror = function(error) {
            console.log(error);
        };
    } catch (error) { 
        console.log("Ошибка", error)
        router.push(`/home`)
    } 
}

router.beforeEach(()=>{
    if (socket != undefined){
        socket.close()
    }
})

// 00 - update player
// 01 - ready | game started
// 02 - turn started
// 03 - player move
// 04 - player bought
// 05 - get all info
// 06 - get quiz
// 07 - end of turn
// 08 - get chance


const handleFunc = (data) =>{
    console.log(data)
    const type = data.substring(0,2)
    let req
    switch (type) {
        case '00':
            req = JSON.parse(data.substring(2))
            loadPlayer(req)
            break;
        case '01':
            console.log('start the game')
            break;
        case '02':
            if (data.substring(2)=='no'){
                return
            }
            req = JSON.parse(data.substring(2))
            startTheTurn(req)
            break;
        case '03':
            req = JSON.parse(data.substring(2))
            movePlayer(req)
            break;
        case '04':
            // TODO: make animation
            quiz.value = null
            break;
        case '05':
            req = JSON.parse(data.substring(2))
            loadRoom(req)
            break
        case '06':
            req = JSON.parse(data.substring(2))
            quiz.value = req
            break
        default:
            break;
    }
}

const loadRoom = (req)=>{
    players.value = req
    for (let player of players.value){
        if (player.username == username.value){
            playerProfile = player
        }
        player.status = 'waiting'
        for (let company of player.companies){
            map.get(company.substring(0,company.length-1)).owner = player.username
            map.get(company.substring(0,company.length-1)).progsCount = Number(company[company.length-1])
        }
    }
}

const startTheTurn = (req)=>{
    let ps = JSON.parse(JSON.stringify(players.value))
    if (activePlayer.value != undefined){
        activePlayer.value.status = 'waiting'
        clearInterval(timerInterval)
    }
            
    for (let i = 0; i<ps.length; i++){
        if (ps[i].username== req.username){
            players.value[i].status = 'active'
            activePlayer.value = players.value[i]
            break
        }
    }
    endOfTurn.value = Date.parse(req.end_of_turn)
}

const move = ()=>{
    console.log('move')
    socket.send('03')   
}
const movePlayer = (req)=>{
    activePlayer.value.position = req.position
}

const loadPlayer = (req)=>{
    for (let player of players.value){
        if (player.username == req.username){
            player.money = req.money
            player.companies = req.companies
            player.position = req.position
            for (let company of player.companies){
                map.get(company.substring(0,company.length-1)).owner = player.username
                map.get(company.substring(0,company.length-1)).progsCount = Number(company[company.length-1])
            }
        }
    }
}

</script>

<style scoped>

.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}

.quiz-wrapper{
    position: absolute;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 100vh;
    backdrop-filter: blur(10px);
    z-index: 100;
    top: 0;
    left: 0;
}

.join-room-modal-wrapper{
    top: 0;
    left: 0;
    position: absolute;
    width: 100vw;
    height: 100vh;
    z-index: 10;
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 15px;
    background-color: var(--black);
    transition: 0.2s;
}
.join-room-modal{
    position: relative;
    display: flex;
    background: var(--black);
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 15px;
    border-radius: 15px;
    border: 3px solid white;
    padding: 25px;
    box-shadow: 0px 0px 10px 4px rgba(255, 255, 255, 0.25);
}
.join-room-modal>span>p{
    font-size: 14px;
}
.join-room-modal>span>input{
    margin-top: 5px;
    width: 250px;
    background: transparent;
    border: 3px solid white;
    border-radius: 15px;
    padding: 5px;
}
.join-room-modal>button{
    position: inherit;
    width: 100%;
    padding: 5px;
}
button{
    position: absolute;
    right: 0;
    padding: 5px;
    border-radius: 15px;
    border: 3px solid #fff;
    background: transparent;
    background-color: var(--black);
    transition: ease-in-out 0.2s;
}
.join-room-modal>button>img{
    height: 100%;
}
.join-room-modal>button:hover{
    transform: scale(1.1);
    transition: ease-in-out 0.2s;
}
#close-modal{
    background: transparent;
    border: none;
    height: 15px;
    position: absolute;
    top: 0;
    right: 0;
    width: 25px;
    height: 25px;
    transform: rotate(45deg);
}
#close-modal:hover{
    transform: scale(1.2) rotate(45deg);
}

.room{
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
}
.row{
    display: flex;
    width: 100%;
    justify-content: space-around;
}
.column{
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    width: 100%;
    height: 100%;
}

</style>