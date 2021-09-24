<template>
  <div class="notif" v-if="msg.length > 0">
    <div class="notif__msg" v-for="message in msg" :key="message.index">
      <div class="notif__msg-icon">
        <i class="fa fa-info-circle"></i>
      </div>
      <div class="notif__msg-content">{{ message.msg }}</div>
      <div class="notif__msg-remove">
        <i class="fa fa-remove"></i>
      </div>
    </div>
  </div>
</template>
<script>
export default {}
</script>
<style lang="scss"></style>

<script>
import { mapState } from 'vuex'

export default {
  computed: {
    ...mapState({
      msg: state => state.auth.msg,
    }),
  },
}
</script>

<style lang="scss" scoped>
@keyframes toleft {
  from {
    right: -30%;
    bottom: 10vh;
  }
  to {
    right: 1%;
    bottom: 10vh;
  }
}
.notif {
  z-index: 2;
  position: fixed;
  right: 1%;
  bottom: 10vh;
  animation: toleft 0.2s ease-in;
  i {
    color: white;
  }
  .notif__msg-icon {
    display: flex;
    justify-content: center;
    i.fa-info-circle {
      font-size: 20px;
    }
  }
  .notif__msg {
    border-radius: 8px;
    padding: 10px 0;
    min-height: 70px;
    width: 400px;
    background-color: #f998a0;
    display: grid;
    grid-template-columns: 60px 1fr 60px;
    align-items: center;
    color: white;
  }
  .notif__msg-remove {
    display: flex;
    justify-content: center;
    i.fa-remove {
      cursor: pointer;
      padding: 5px 7px;
      background-color: #f998a0;
      border-radius: 50%;
    }
  }
}
</style>
