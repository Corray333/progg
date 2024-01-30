<template>
    <div class="rooms">
        <div class="create-room-modal-wrapper" v-show="showModal">
            <div class="create-room-modal">
                <h2>Create room</h2>
                <span>
                    <p>Username: </p>
                    <input type="text" v-model="username" autocomplite="off">
                </span>
                <span>
                    <p>Name: </p>
                    <input type="text" v-model="roomName" autocomplite="off">
                </span>
                <span>
                    <p>Password: </p>
                    <input type="password" v-model="password" autocomplite="off">
                </span>
                <button @click="createRoom()">Create room</button>
                <button id="close-modal" @click="showModal = false"><img src="../assets/icons/plus.png" alt=""></button>
            </div>
        </div>
        
        <h1>Rooms</h1>
        <form class="search-bar" @submit.prevent="find()">
            <input type="text" v-model="search">
            <button @click="find()"><img src="../assets/icons/search.png" alt=""></button>
            <button @click="showModal = true"><img src="../assets/icons/plus.png" alt=""></button>
        </form>
        <div class="rooms-list">
            <div class="room" v-for="(room, id) in rooms" :key="id">
                <p>{{room}}</p>
                <button @click="joinRoom(room)"><img src="../assets/icons/play.png" alt=""></button>
            </div>
        </div>  
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()

const rooms = ref([])
const search = ref('')
const number = ref(0)
const showModal = ref(false)

const username = ref('')
const password = ref('')
const roomName = ref('')

const find = async () => {
    try { 
        const resp = await axios.get(`http://localhost:3000/rooms?number=${number.value}&name=${search.value}`) 
        rooms.value = resp.data.rooms
    } catch (error) { 
        alert("Ошибка", error)
    } 
}

const createRoom = async ()=>{
    try { 
        let req = {
            room_name:roomName.value,
            player_name:username.value,
            password:password.value
        }
        await axios.post(`http://localhost:3000/rooms`, req) 
        router.push(`/rooms/${roomName.value}`)
    } catch (error) { 
        alert("Ошибка", error)
    } 
}


onMounted(()=>{
    find()
})


</script>

<style scoped>

.rooms{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
    gap:15px;
}
.search-bar{
    width: 300px;
    display: flex;
    position: relative;
}
.search-bar>input{
    width: 100%;
    font-size: 16px;
    padding: 5px;
    border-radius: 16px;
    border: 3px solid #fff;
    background: transparent;
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
button>img{
    height: 100%;
}
button:hover{
    transform: scale(1.1);
    transition: ease-in-out 0.2s;
}
.search-bar>button:last-child{
    right: -50px;
    padding: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
}



.room {
    display: flex;
    align-items: center;
    width: 300px;
    background: transparent;
    position: relative;
    margin-bottom: 5px;
}
.room>p{
    padding: 5px;
    border: 3px solid #fff;
    font-size: 16px;
    border-radius: 15px;
    width: 100%;
}
.room>a{
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
.room>a>img{
    height: 100%;
}
.room>a:hover{
    transform: scale(1.1);
    transition: ease-in-out 0.2s;
}


.create-room-modal-wrapper{
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
    backdrop-filter: blur(15px);
    transition: 0.2s;
}
.create-room-modal{
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
span>p{
    font-size: 14px;
}
span>input{
    margin-top: 5px;
    width: 250px;
    background: transparent;
    border: 3px solid white;
    border-radius: 15px;
    padding: 5px;
}
.create-room-modal>button{
    position: inherit;
    width: 100%;
    padding: 5px;
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


</style>