<template>
  <NavBar/>
  <router-view/>
</template>

<script setup>
import NavBar from '@/components/NavBar.vue'
</script>

<style>
:root{
  --black: #0F0D17;
  --green: #2CEC3F;
  --blue: #2C34EC;
  --red: #EC2C2C;
  --purple: #A32CEC;
  --yellow: #ECB72C;
  --orange: #F17732;
  --light-blue: #54B7E1;
}
*{
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: "Ubuntu Mono";
  color: white;
}
body{
  background-color: var(--black);
  width: 100%;
}
#app{
  display: flex;
  width: 100%;
  gap: 25px;
  padding: 25px;
  min-height: 100vh;
}
*:focus{
  outline: none;
}
</style>