const appLogin = createApp({
  delimiters: ['@{', '}'],
  data(){
    return {
      user: {username: '', password: ''},
    }
  }, 
  methods: {
    loginUser(){
      if (this.user.username == ''){
        console.log('Chưa nhập tên đăng nhập');
      }else{
        axios.post('login', {username: this.user.username, password: this.user.password})
        .then(response => {
              var data = response.data;
              localStorage.setItem('refresh_token', data.refresh_token)
              window.location.href = '/invoice';          
            })
        .catch(error => {
          console.log(error.data.Error);
          });      
      }
    },
    checkForEnter(event){
      if (event.key == "Enter") {
        this.loginUser();
      }
    },
  }
})
appLogin.mount('#root')