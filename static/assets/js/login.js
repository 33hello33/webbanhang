const {createApp} =  Vue;
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
        alert('Chưa nhập tên đăng nhập');
      }else{
        axios.post('login', {username: this.user.username, password: this.user.password}).then(response => {
            if(response.status == 200){
              var data = response.data;
              $cookies.set('token', data.token);
              localStorage.setItem('refresh_token', data.refresh_token)
              window.location.href = '/';          
            }else{
              console.log(response.data)
            } 
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