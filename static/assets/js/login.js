var Vue = new Vue({
  el: '#root',
  delimiters: ['@{', '}'],
  data: {
    showError: false,
    enableEdit: false,
    user: {username: '', password: ''},
  },  methods: {
    loginUser(){
      if (this.user.username == ''){
        this.showError = true;
      }else{
        this.showError = false;
          this.$http.post('login', {username: this.user.username, password: this.user.password}).then(response => {
            var data = response.body;
            $cookies.set('token', data.token);
            localStorage.setItem('refresh_token', data.refresh_token)
            window.location.href = '/';          
          });      
        }
    },
    checkForEnter(event){
      if (event.key == "Enter") {
        this.loginUser();
      }
    },
  }
});