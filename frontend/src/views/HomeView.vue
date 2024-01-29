<template>
    <div class="home">
        <img src="../assets/icons/logo-white.png" alt="">
        <h1>Монополия для программистов😎</h1>
        <hr>
        <article>
            <h2>Описание</h2>
            <p></p>
        </article>
        <hr>
        <article>
            <h2>Правила</h2>
            <p></p>
        </article>
    </div>
</template>

<script setup>

</script>

<style scoped>

.home{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
}
.home>img{
    width: 230px;
}
.home>h1{
    margin-top: 5px;
    font-size: 14px;
    font-weight: normal;
}

hr{
    width: 300px;
}

</style>