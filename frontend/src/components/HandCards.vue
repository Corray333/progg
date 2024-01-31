<template>
    <div class="hand">
        <div class="card" v-for="(card, i) of cards" :key="i">
            <p>Шанс</p>
            <p>{{ Cards.get(card) }}</p>
        </div>
    </div>
</template>

<script setup>
import {ref} from 'vue'

const cards = ref([1,0,5])

const Cards = new Map()
Cards.set(0, "Отправляйтесь в тюрьму.")
Cards.set(1, "Получите 200к.")
Cards.set(2, "Выплатите 200к.")
Cards.set(3, "Отправляйтесь на старт.")
Cards.set(4, "Вернитесь на три поля назад.")
Cards.set(5, "Отправляйтесь на три поля вперед.")

</script>

<style scoped>

.hand{
    display: flex;
}
.card{
    height: 300px;
    width: 225px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    border: 3px solid white;
    border-radius: 15px;
    margin-left: -150px;
    background-color: var(--black);
    transition: ease-in-out 0.2s;
}
.card:hover{
    z-index: 10;
    transform: scale(1.05) translateY(-20px);
    box-shadow: 0 0 20px 0 rgba(255,255,255,0.5);
    transition: ease-in-out 0.2s;
}
.card>p:first-child{
    font-size: 24px;
    font-weight: bold;
}
p{
    text-align: center;
}

</style>