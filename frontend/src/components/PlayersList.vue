<template>
    <div class="players">
        <h2>Players</h2>
        <div class="player" v-for="(player, i) of props.players" :key="i">
            <img :src="`/src/assets/avatars/${player.avatar}.png`" alt="">
            <div class="player-info">
                <p>{{ player.username }}</p>
                <p>{{ player.money }}k</p>
            </div>
        </div>
        <button @click="$emit('ready')">Готов</button>
    </div>
</template>

<script setup>
const props = defineProps(['players'])




</script>

<style scoped>

.players{
    display: flex;
    flex-direction: column;
    align-items: center;
    gap:15px;
}
.player{
    display: flex;
    align-items: center;
    gap: 5px;
    width: 250px;
    border: 3px solid white;
    border-radius: 15px;
    padding: 10px;
}
.player>img{
    width: 50px;
    border: 3px solid white;
    border-radius: 50%;
}
.players>button{
    width: 100%;
}

</style>