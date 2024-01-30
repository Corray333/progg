<template>
    <div class="map">
        <img src="../assets/icons/arrow.png" alt="" id="arrow">
        <div class="row row1">
            <div class="map-tile">
                <img src="../assets/icons/prison.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'megafon')">
                <img src="../assets/icons/megafon.png" alt="">
                <span class="green-line"></span>
            </div>
            <div class="map-tile">
                <p>?</p>
                <p>Шанс</p>
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'kaspersky')">
                <img src="../assets/icons/kaspersky.png" alt="">
                <span class="green-line"></span>
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'sber')">
                <img src="../assets/icons/sber.png" alt="">
                <span class="green-line"></span>
            </div>
            <div class="map-tile">
                <p>?</p>
                <p>Шанс</p>
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'astra')">
                <img src="../assets/icons/astra.png" alt="">
                <span class="white-line"></span>
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'geekbrains')">
                <img src="../assets/icons/geekbrains.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'rostech')">
                <img src="../assets/icons/rostech.png" alt="">
                <span class="white-line"></span>
            </div>
            <div class="map-tile">
                <img src="../assets/icons/sleep.png" alt="">
            </div>
        </div>
        <div class="row row2">
            <div class="column column1">
                <div class="map-tile" @click="$emit('tile-pick', 'tensor')">
                    <img src="../assets/icons/tensor.png" alt="">
                    <span class="orange-line"></span>
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'finam')">
                    <img src="../assets/icons/finam.png" alt="">
                    <span class="orange-line"></span>
                </div>
                <div class="map-tile">
                    <p>?</p>
                    <p>Шанс</p>
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'europlan')">
                    <img src="../assets/icons/europlan.png" alt="">
                    <span class="orange-line"></span>
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'hh')">
                    <img src="../assets/icons/hh.png" alt="">
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'yota')">
                    <img src="../assets/icons/yota.png" alt="">
                    <span class="light-blue-line"></span>
                </div>
                <div class="map-tile">
                    <p>?</p>
                    <p>Шанс</p>
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'skyeng')">
                    <img src="../assets/icons/skyeng.png" alt="">
                    <span class="light-blue-line"></span>
                </div>
            </div>
            <div class="chance-deck">
                <div class="chance-card"><p>Шанс</p></div>
                <div class="chance-card"><p>Шанс</p></div>
            </div>
            <div class="column column2">
                <div class="map-tile" @click="$emit('tile-pick', '1c')">
                    <span class="yellow-line"></span>
                    <img src="../assets/icons/1c.png" alt="">
                </div>
                <div class="map-tile">
                    <p>Шанс</p>
                    <p>?</p>
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'beeline')">
                    <span class="yellow-line"></span>
                    <img src="../assets/icons/beeline.png" alt="">
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'tinkoff')">
                    <span class="yellow-line"></span>
                    <img src="../assets/icons/tinkoff.png" alt="">
                </div>
                <div class="map-tile">
                    <p>Шанс</p>
                    <p>?</p>
                </div>

                <div class="map-tile" @click="$emit('tile-pick', 'rostelekom')">
                    <span class="purple-line"></span>
                    <img src="../assets/icons/rostelekom.png" alt="">
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'wot')">
                    <img src="../assets/icons/wot.png" alt="">
                </div>
                <div class="map-tile" @click="$emit('tile-pick', 'netbynet')">
                    <span class="purple-line"></span>
                    <img src="../assets/icons/netbynet.png" alt="">
                </div>
            </div>
        </div>
        <div class="row row3">
            <div class="map-tile" @click="$emit('tile-pick', 'reward')">
                <img src="../assets/icons/reward.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'vk')">
                <span class="blue-line"></span>
                <img src="../assets/icons/vk.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'intellij')">
                <img src="../assets/icons/intellij.png" alt="">                
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'ozon')">
                <span class="blue-line"></span>
                <img src="../assets/icons/ozon.png" alt="">
            </div>
            <div class="map-tile">
                <p>Шанс</p>
                <p>?</p>
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'alpha')">
                <span class="red-line"></span>
                <img src="../assets/icons/alpha.png" alt="">
            </div>
            <div class="map-tile">
                <p>Шанс</p>
                <p>?</p>
            </div>
            <div class="map-tile"  @click="$emit('tile-pick', 'yandex')">
                <span class="red-line"></span>
                <img src="../assets/icons/yandex.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'magnit')">
                <span class="red-line"></span>
                <img src="../assets/icons/magnit.png" alt="">
            </div>
            <div class="map-tile" @click="$emit('tile-pick', 'arest')">
                <img src="../assets/icons/arest.png" alt="">
            </div>
        </div>
    </div>
</template>

<script setup>


</script>

<style scoped>

#arrow{
    position: absolute;
    bottom: 25px;
    left: -50px;
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

</style>