<template>
  <div class="container">
    <div class="columns">
      <div class="column is-12">
        <h1 class="heading is-large">Lass und dein MeiSo Verbinden!</h1>
        <p>blablabla</p>
      </div>
    </div>
    <div v-if="!connected" class="columns">
      <div class="column">
        <button class="button is-primary" @click="connect">Verbinden</button>
      </div>
    </div>
    <div v-else class="columns is-multiline">
      <div class="column is-12">
        <div class="field">
          <label class="label">Wlan Name (SSID)</label>
          <div class="control">
            <input class="input" type="text" placeholder="Wlan Name (SSID)" v-model="ssid">
          </div>
        </div>
      </div>
      <div class="column is-12">
        <div class="field">
          <label class="label">Passwort</label>
          <div class="control">
            <input class="input" type="password" placeholder="Passwort" v-model="password">
          </div>
        </div>
      </div>
      <div class="column">
        <button class="button is-primary" @click="handleSubmit">Speichern</button>
      </div>
    </div>
    <FlowerLoader v-if="connecting && !connected" />
  </div>
</template>

<script>
import axios from 'axios'
import FlowerLoader from '@/components/loader/FlowerLoader'
export default {
  name: 'WifiSetup',
  components: { FlowerLoader },
  data () {
    return {
      connected: false,
      connecting: false,
      ssid: null,
      password: null
    }
  },
  methods: {
    connect () {
      this.connecting = true
      this.healthCheck()
      let intervalId = setInterval(() => {
        if (!this.connected) {
          this.healthCheck()
        } else {
          clearInterval(intervalId)
          intervalId = null
        }
      }, 10000)
    },
    healthCheck: async function () {
      await axios.get('http://192.168.4.1/health', {
        timeout: 10000
      })
        .then((resp) => {
          if (resp.status === 200) {
            this.connected = true
            this.connecting = false
          }
        })
        .catch(() => {
          this.connected = false
        })
    },
    handleSubmit () {
      const plain = `${this.ssid}\n${this.password}\n`
      console.log(plain)
      axios.post('http://192.168.4.1/wifi', plain, {
        headers: {
          'Content-Type': 'text/plain'
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
