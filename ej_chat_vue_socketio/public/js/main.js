const socket = io();

new Vue({
  el: '#chat-app',
  created() {
    socket.on("mensaje chat", (message) => {
      console.log(message)
      this.messages.push({
        text: message,
        date: new Date().toLocaleString()
      })
    })
  },
  data: {
    message: '',
    messages: []
  },
  methods: {
    sendMessage() {
      socket.emit("mensaje chat", this.message)
      this.message = "";
    }
  }
})