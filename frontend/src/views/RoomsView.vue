<template>
    <div class="rooms">
        <div class="search-bar">
            <input type="text" v-model="search">
            <button @click="find()"><img src="../assets/icons/search.png" alt=""></button>
        </div>
        <div class="rooms-list">
            <div class="room" v-for="(room, id) in rooms" :key="id">
                <p>{{room}}</p>
                <button><img src="../assets/icons/play.png" alt=""></button>
            </div>
        </div>  
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'


const rooms = ref([])
const search = ref('')
const number = ref(0)

const find = async () => {
    try { 
        const resp = await axios.get(`http://localhost:3000/rooms?number=${number.value}&name=${search.value}`) 
        rooms.value = resp.data.rooms
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

.room {
    font-size: 16px;
    display: flex;
    align-items: center;
    width: 300px;
    border-radius: 15px;
    border: 3px solid #fff;
    background: transparent;
    position: relative;
}
.room>p{
    padding: 5px;

}

</style>