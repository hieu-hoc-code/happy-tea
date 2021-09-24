<template>
  <div class="notif">
    <div class="notif__msg">
      <i class="fa fa-info-circle"></i>
      <div class="notif__msg-content">Success: cac kieuc con da dieu</div>
      <i class="fa fa-remove"></i>
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
  from{
    right: -30%;
    bottom: 10vh;
  }
  to{
    right: 1%;
    bottom: 10vh;
  }
}
.notif {
  z-index: 2;
  position: fixed;
  height: 80px;
  animation: toleft 2s infinite;
  i{
    color: white;
  }
  i.fa-info-circle{
    font-size: 20px;
  }
  .notif__msg {
    border-radius: 8px;
    min-height: 80px;
    width: 400px;
    background: #3e8ed0;
    display: flex;
    justify-content: space-around;
    align-items: center;
    color: white;
  }
  i.fa-remove{
    cursor: pointer;
    padding: 5px 7px;
    background-color: #3373a8;
    border-radius: 50%;
  }
}
</style>
