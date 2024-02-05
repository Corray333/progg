<template>
    <div class="map" ref="map">
        <img src="../assets/icons/arrow.png" class="clickable" alt="" id="arrow" @click="$emit('move')" v-if="props.playerProfile != undefined && props.playerProfile.status == 'active' & props.playerProfile.already_walked==false">
        <img src="../assets/icons/arrow.png" alt="" id="arrow" v-if="props.playerProfile != undefined && props.playerProfile.status == 'waiting'">
        <div class="players" style="display:none;" ref="playersBlock">
            <div class="player" v-for="(player, i) of props.players" :key="i" :id="player.username"> 
                <img :src="`/src/assets/avatars/${player.avatar}.png`" alt="">
                <div class="player-line"></div>
            </div>
        </div>
        <div class="row row1">
            <div class="map-tile" id="i10">
                <img src="../assets/icons/prison.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'megafon')" id="i11">
                <img src="../assets/icons/megafon.png" alt="">
                <span class="green-line"></span>
            </div>
            <div class="map-tile" id="i12">
                <p>?</p>
                <p>Шанс</p>
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'kaspersky')" id="i13">
                <img src="../assets/icons/kaspersky.png" alt="" id="">
                <span class="green-line"></span>
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'sber')" id="i14">
                <img src="../assets/icons/sber.png" alt="">
                <span class="green-line"></span>
            </div>
            <div class="map-tile" id="i15">
                <p>?</p>
                <p>Шанс</p>
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'astra')" id="i16">
                <img src="../assets/icons/astra.png" alt="">
                <span class="white-line"></span>
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'geekbrains')" id="i17">
                <img src="../assets/icons/geekbrains.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'rostech')" id="i18">
                <img src="../assets/icons/rostech.png" alt="">
                <span class="white-line"></span>
            </div>
            <div class="map-tile" id="i19">
                <img src="../assets/icons/sleep.png" alt="">
            </div>
        </div>
        <div class="row row2">
            <div class="column column1">
                <div class="map-tile" @click="$emit('tile-pick', 'finam')" id="i9">
                    <img src="../assets/icons/finam.png" alt="">
                    <span class="orange-line"></span>
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'tensor')" id="i8">
                    <img src="../assets/icons/tensor.png" alt="">
                    <span class="orange-line"></span>
                </div>
                <div class="map-tile" id="i7">
                    <p>?</p>
                    <p>Шанс</p>
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'europlan')" id="i6">
                    <img src="../assets/icons/europlan.png" alt="">
                    <span class="orange-line"></span>
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'hh')" id="i5">
                    <img src="../assets/icons/hh.png" alt="">
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'yota')" id="i4">
                    <img src="../assets/icons/yota.png" alt="">
                    <span class="light-blue-line"></span>
                </div>
                <div class="map-tile" id="i3">
                    <p>?</p>
                    <p>Шанс</p>
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'skyeng')" id="i2">
                    <img src="../assets/icons/skyeng.png" alt="">
                    <span class="light-blue-line"></span>
                </div>
            </div>
            <div class="chance-deck" @click="placePlayers()">
                <div class="chance-card"><p>Шанс</p></div>
                <div class="chance-card"><p>Шанс</p></div>
            </div>
            <div class="column column2">
                <div class="map-tile" @click="$emit('tile-pick', '1c')" id="i20">
                    <span class="yellow-line"></span>
                    <img src="../assets/icons/1c.png" alt="">
                </div>
                <div class="map-tile" id="i21">
                    <p>Шанс</p>
                    <p>?</p>
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'beeline')" id="i22">
                    <span class="yellow-line"></span>
                    <img src="../assets/icons/beeline.png" alt="">
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'tinkoff')" id="i23">
                    <span class="yellow-line"></span>
                    <img src="../assets/icons/tinkoff.png" alt="">
                </div>
                <div class="map-tile" id="i24">
                    <p>Шанс</p>
                    <p>?</p>
                </div>

                <div class="map-tile" @click="$emit('tile-pick', 'rostelekom')" id="i25">
                    <span class="purple-line"></span>
                    <img src="../assets/icons/rostelekom.png" alt="">
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'wot')" id="i26">
                    <img src="../assets/icons/wot.png" alt="">
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'netbynet')" id="i27">
                    <span class="purple-line"></span>
                    <img src="../assets/icons/netbynet.png" alt="">
                </div>
            </div>
        </div>
        <div class="row row3">
            <div class="map-tile" @click="$emit('tile-pick', 'reward')" id="i1">
                <img src="../assets/icons/reward.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'vk')" id="i36">
                <span class="blue-line"></span>
                <img src="../assets/icons/vk.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'intellij')" id="i35">
                <img src="../assets/icons/intellij.png" alt="">                
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'ozon')"  id="i34">
                <span class="blue-line"></span>
                <img src="../assets/icons/ozon.png" alt="">
            </div>
            <div class="map-tile" id="i33">
                <p>Шанс</p>
                <p>?</p>
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'alpha')" id="i32">
                <span class="red-line"></span>
                <img src="../assets/icons/alpha.png" alt="">
            </div>
            <div class="map-tile" id="i31">
                <p>Шанс</p>
                <p>?</p>
            </div>
            <div class="map-tile"  @click="$emit('tile-pick', 'yandex')" id="i30">
                <span class="red-line"></span>
                <img src="../assets/icons/yandex.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'magnit')" id="i29">
                <span class="red-line"></span>
                <img src="../assets/icons/magnit.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'arest')" id="i28">
                <img src="../assets/icons/arest.png" alt="">
            </div>
        </div>
    </div>
</template>

<script setup>
import { defineProps, ref, defineExpose } from 'vue'
const props = defineProps(['players', 'endOfTurn', 'playerProfile'])


const map = ref(null)
const playersBlock = ref(null)


const placePlayers = ()=>{
    if (props.players == undefined) return
    playersBlock.value.style.display = 'block'
    for (let player of props.players){
        let p = map.value.querySelector(`#${player.username}`)
        let tile = map.value.querySelector(`#i${player.position}`)
        switch (true){
            case player.position == 1:
                p.style.top = `${tile.offsetTop-75}px`
                p.style.left = `${tile.offsetWidth}px`
                p.style.transform = `rotate(45deg)`
                break
            case player.position >= 2 && player.position <= 9:
                p.style.top = `${tile.offsetTop}px`
                p.style.left = `${tile.offsetWidth}px`
                p.style.transform = `rotate(90deg)`
                break
            case player.position == 10:
                p.style.top = `${tile.offsetHeight}px`
                p.style.left = `${tile.offsetWidth}px`
                p.style.transform = `rotate(135deg)`
                break
            case player.position >= 11  && player.position <= 18:
                p.style.top = `${tile.offsetHeight}px`
                p.style.left = `${tile.offsetLeft}px`
                p.style.transform = `rotate(180deg)`
                break
            case player.position == 19:
                p.style.top = `${tile.offsetHeight}px`
                p.style.left = `${tile.offsetLeft-75}px`
                p.style.transform = `rotate(225deg)`
                break
            case player.position >= 20  && player.position <= 27:
                p.style.top = `${tile.offsetTop}px`
                p.style.left = `${tile.offsetLeft-75}px`
                p.style.transform = `rotate(270deg)`
                break
            case player.position == 28:
                p.style.top = `${tile.offsetTop-75}px`
                p.style.left = `${tile.offsetLeft-75}px`
                p.style.transform = `rotate(315deg)`
                break
            case player.position >= 29  && player.position <= 36:
                p.style.top = `${tile.offsetTop-75}px`
                p.style.left = `${tile.offsetLeft}px`
                p.style.transform = `rotate(360deg)`
                break
            default:
                break
        }

    }
}

defineExpose({
  placePlayers
})




</script>

<style scoped>

#arrow{
    position: absolute;
    bottom: 25px;
    left: -50px;
}
.clickable{
    cursor: pointer;
    transition: 0.2s;
}
.clickable:hover{
    transform: scale(1.1);
    transition: 0.2s;
}
.red-line{
    background-color: var(--red);
}
.purple-line{
    background-color: var(--purple);
}
.yellow-line{
    background-color: var(--yellow);
}
.green-line{
    background-color: var(--green);
}
.blue-line{
    background-color: var(--blue);
}
.orange-line{
    background-color: var(--orange);
}
.white-line{
    background-color: white;
}
.light-blue-line{
    background-color: var(--light-blue);
}

.map{
    width: fit-content;
    position: relative;
}
.map-tile{
    width: 75px;
    height: 100px;
    border: 3px solid white;
    margin: 0 -1px;
}
.row{
    display: flex;
    justify-content: space-between;
}
.row>.map-tile:first-child, .row>.map-tile:last-child{
    width: 100px;
    margin: 0;
    margin-bottom: -1px;
}
.row>.map-tile:first-child{
    margin-right: -1px;
}
.row>.map-tile:last-child{
    margin-left: -1px;
}
.row1>.map-tile:first-child{
    border-radius: 15px 0 0 0;
}
.row1>.map-tile:last-child{
    border-radius: 0 15px 0 0;
}
.row3>.map-tile:last-child{
    border-radius: 0 0 15px 0;
}
.row3>.map-tile:first-child{
    border-radius: 0 0 0 15px;
}
.row1{
    margin-bottom: -1px;
}
.row3{
    margin-top: -1px;
}
.column{
    display: flex;
    flex-direction: column;
}
.column>.map-tile{
    width: 100px;
    height: 75px;
    margin: -1px 0;
}

.row>.map-tile{
    flex-direction: column;
}
.map-tile{
    display: flex;
    justify-content: space-around;
    align-items: center;
    background-color: var(--black);
    transition: 0.2s;
}
.map-tile:hover{
    transform: scale(1.1);
    box-shadow: 0 0 10px 0px white;
    transition: 0.2s;
    border-radius: 5px;
    cursor: pointer;
}
.map-tile>span {
    width: 20px;
    height: 5px;
    border-radius: 99px;
    margin: 0 20px;
}
.row2{
    align-items: center;
}
.row1>.map-tile>p:first-child, 
.column1>.map-tile>p:first-child{
    font-size: 64px;
    font-weight: bold;
    line-height: 87%;
}
.row3>.map-tile>p:last-child, 
.column2>.map-tile>p:last-child{
    font-size: 64px;
    font-weight: bold;
    line-height: 87%;
}

.column>.map-tile{
    flex-direction: row;
}
.column>.map-tile>span{
    height: 20px;
    width: 5px;
    margin: 20px 0;
}
.column1>.map-tile>p:last-child{
    transform: rotate(90deg);
}
.column2>.map-tile>p:first-child{
    transform: rotate(-90deg);
}
.column1>.map-tile>p:first-child{
    padding-left: 20px;
}
.column2>.map-tile>p:last-child{
    padding-right: 20px;
}

.chance-deck{
    position: relative;
    width: 300px;
    height: 225px;
}
.chance-card{
    width: 300px;
    height: 225px;
    border: 3px solid white;
    background-color: var(--black);
    border-radius: 15px;
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    transition: ease-in-out 0.2s;
}
.chance-card>p{
    font-size: 24px;
    font-weight: bold;
}
.chance-card:last-child:hover{
    transform: scale(1.05) translateY(-25px);
    box-shadow: 0 0 20px 0px rgba(255, 255, 255, 0.5);
    transition: ease-in-out 0.2s;
    cursor: pointer;
}

.player{
    width: 75px;
    height: 75px;
    position: absolute;
    transition: ease-in-out 1s;
    padding: 10px;
}
.player>img{
    width: 100%;
    height: 100%;
    border-radius: 99px;
    border: 3px solid white;
    z-index: 10;
    position: relative;
}
.player-line{
    top: 50%;
    left: 50%;
    height: 50%;
    position: absolute;
    z-index: 1;
    width: 3px;
    background-color: white;
    border-radius: 3px;
}


</style>