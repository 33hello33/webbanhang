var Vue = new Vue({
  el: '#root',
  delimiters: ['&{', '}'],
  data: {
    user: {name: '', username: '', password: '', email: ''},
  },
  
  methods: {
    registerUser(){

          this.$http.post('register',{name: this.user.name, username: this.user.username, password: this.user.password, email: this.user.email}).then(response => {
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
