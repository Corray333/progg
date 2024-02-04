<template>
    <div class="quiz-modal">
        <h2>{{ props.quiz.question }}</h2>
        <ul>
            <li v-if="props.playerProfile != undefined && props.activePlayer != undefined && props.activePlayer.username == props.playerProfile.username" v-for="(variant, i) of props.quiz.variants" :key="i" @click="$emit('answer', variant)">{{ i }}. {{ variant }}</li>
            <li class="inactive" v-if="props.playerProfile != undefined && props.activePlayer != undefined && props.activePlayer.username != props.playerProfile.username" v-for="(variant, i) of props.quiz.variants" :key="i" >{{ i }}. {{ variant }}</li>
        </ul>
    </div>
</template>

<script setup>

const props = defineProps(['quiz', 'playerProfile', 'activePlayer'])

</script>

<style scoped>

.quiz-modal {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    border: 3px solid white;
    border-radius: 15px;
    padding: 25px;
    background-color: var(--black);
    gap: 10px;
}
ul{
    display: flex;
    flex-direction: column;
    gap: 10px;
}
ul>li{
    border: 3px solid white;
    padding: 5px;
    border-radius: 15px;
    transition: ease-in-out 0.3s;
    list-style: none;
    cursor: pointer;
}
ul>li:hover{
    transform: scale(1.1);
    transition: ease-in-out 0.2s;
}
.inactive:hover{
    transform: none;
}

 </style>