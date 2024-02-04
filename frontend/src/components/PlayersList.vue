<template>
    <div class="players">
        <h2>Players</h2>
        <div class="player" :class="player.status" v-for="(player, i) of props.players" :key="i">
            <img :src="`/src/assets/avatars/${player.avatar}.png`" alt="">
            <div class="player-info">
                <p>{{ player.username }}</p>
                <p>{{ player.money }}k</p>
            </div>
            <p v-if="player.status == 'active'" class="timer">{{ turnTimer }}c</p>
            <button @click="$emit('stopTheTurn')" v-if="props.playerProfile != undefined && props.activePlayer != undefined && props.playerProfile.username == props.activePlayer.username && props.playerProfile.username == player.username">
                <img src="/src/assets/icons/next.png" alt="">
            </button>
        </div>
        <button @click="$emit('ready')" v-if="playerProfile != undefined && playerProfile.ready==false">Готов</button>
    </div>
</template>

<script setup>
const props = defineProps(['players', 'turnTimer', 'playerProfile', 'activePlayer'])






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

.waiting{
    border-color: white;
}
.active {
    border-color: var(--yellow);
}

.timer{
    justify-self: flex-end;
}

.player>button>img{
    width: 20px;
}

</style>