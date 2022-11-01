appRegister = createApp({
  delimiters: ['&{', '}'],
  data() {
    return {
      user: {name: '', username: '', password: '', email: ''},
    }
  },
  
  methods: {
    registerUser(){
          axios.post('register',{name: this.user.name, username: this.user.username, password: this.user.password, email: this.user.email}).then(response => {
            window.location.href = "/login"
          });
    },
    checkForEnter(event){
      if (event.key == "Enter") {
        this.registerUser();
      }
    },
  }
});
appRegister.mount("#root")