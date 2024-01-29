<template>
    <navbar>
        <router-link to="/home"><img src="../assets/icons/home.png" alt=""></router-link>
        <router-link to="/rooms"><img src="../assets/icons/play.png" alt=""></router-link>
    </navbar>    
</template>

<script setup>

</script>

<style scoped>

navbar{
    display: flex;
    flex-direction: column;
    gap: 15px;
}
img{
    width: 30px;
    transition:ease-in-out 0.2s;
}
img:hover{
    transform: scale(1.2);
    transition:ease-in-out 0.2s;
}
</style>