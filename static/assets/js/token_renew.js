var Vue = new Vue({
    el: '#root',
    delimiters: ['@{', '}'],
   mounted() { 
    var refresh_token = localStorage.getItem('refresh_token');
    
    this.$http.post('/token/renew', {"refresh_token" : refresh_token}).then( response =>{
      if(response.status == 200){
        token = response.body.token;
        $cookies.set('token', token);
        window.location.href = '/';          
      }
    })
  },
});

