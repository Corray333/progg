<template>
    <div class="room">
        <div class="join-room-modal-wrapper" v-show="showModal">
            <div class="join-room-modal">
                <h2>Join room</h2>
                <span>
                    <p>Username: </p>
                    <input type="text" v-model="username" autocomplite="off">
                </span>
                <span>
                    <p>Password: </p>
                    <input type="password" v-model="password" autocomplite="off">
                </span>
                <button @click="joinRoom()">Join room</button>
                <button id="close-modal" @click="showModal = false"><img src="../assets/icons/plus.png" alt=""></button>
            </div>
        </div>
        <PlayMap @tile-pick="changeCard"/>
        <div class="column">
            <div class="row">
                <InfoCard :pick="pick"/>
                <PlayersList/>
            </div>
            <HandCards/>
        </div>
    </div>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router'
import {ref} from 'vue'
import axios from 'axios'
import PlayMap from '@/components/PlayMap.vue'
import InfoCard from '@/components/InfoCard.vue'
import PlayersList from '@/components/PlayersList.vue'
import HandCards from '@/components/HandCards.vue'

const route = useRoute().params.room
const router = useRouter()
console.log(route)

const showModal = ref(true)
const username = ref('')
const password = ref('')


// const room = ref()

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
        img: 'icons/vk.png'
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
        img: 'icons/ozon.png'
    }
)
map.set('yandex', {
        name:'Yandex',
        description:'Yandex — российская транснациональная компания, владеющая одноимённой системой поиска в Сети, интернет-порталами и службами в нескольких странах. Штаб-квартира — в Москве.',
        price:320,
        color:'red',
        programs:['Yandex','Yandex музыка','Yandex облако'],
        progsCount: 0,
        rentLevels: [28, 250, 450, 1000],
        progPrice: 200,
        owner: 'отсутствует',
        playersHere: [],
        img: 'icons/yandex.png'
    }
)
map.set('alpha', {
        name:'Ozon',
        description:'Ozon — российская интернет-компания, владеющая одноимённым интернет-магазином. Штаб-квартира компании расположена в Москве. Основана в 1998 году Дмитрием Бакшеевым и Александром Шульгином.',
        price:300,
        color:'red',
        programs:['Aplha банк','Alpha страхование','Alpha инвестиции'],
        progsCount: 0,
        rentLevels: [26, 130, 390, 900],
        progPrice: 200,
        owner: 'отсутствует',
        playersHere: [],
        img: 'icons/yandex.png'
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
        img: 'icons/magnit.png'
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
        img: 'icons/rostelekom.png'
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
        img: 'icons/netbynet.png'
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
        img: 'icons/tinkoff.png'
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
        img: 'icons/beeline.png'
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
        img: 'icons/1c.png'
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
        img: 'icons/rostech.png'
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
        img: 'icons/astra.png'
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
        img: 'icons/sber.png'
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
        img: 'icons/kaspersky.png'
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
        img: 'icons/megafon.png'
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
        img: 'icons/finam.png'
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
        img: 'icons/tensor.png'
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
        img: 'icons/europlan.png'
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
        img: 'icons/yota.png'
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
        img: 'icons/skyeng.png'
    }
)
}

const pick = ref()
const changeCard = (query)=>{
    pick.value = map.get(query)
}

const joinRoom = async ()=>{
    try { 
        let req = {
            room_name:route,
            player_name:username.value,
            password:password.value
        }
        await axios.post(`http://localhost:3000/rooms`, req) 
        let socket = new WebSocket(`ws://localhost:3000/rooms/${route}/game`);

        socket.onopen = function() {
            console.log("[open] Connection established");
            console.log("Sending to server");
            socket.send(username.value);
            showModal.value = false
        };

        socket.onmessage = function(event) {
            handleFunc(event.data)
        };

        socket.onclose = function(event) {
            if (event.wasClean) {
                console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
            } else {
                console.log('[close] Connection died');
            }
        };

        socket.onerror = function(error) {
            console.log(error);
        };
    } catch (error) { 
        alert("Ошибка", error)
        router.push(`/home`)
    } 
}

// 00 - exit
// 01 - start the game
// 02 - turn started
// 03 - player moved
// 04 - player bought


const handleFunc = (data) =>{
    const req = JSON.parse(data)
    console.log(req)
}

</script>

<style scoped>

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
.join-room-modalspan>input{
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
    height: 100%;
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
}

</style>