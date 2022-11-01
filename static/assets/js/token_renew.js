appTokenRenew = createApp({
    delimiters: ['@{', '}'],
    beforeMount() { 
      var refresh_token = localStorage.getItem('refresh_token');
      
      axios.post('/token/renew', {"refresh_token" : refresh_token})
      .then( response =>{
        if(response.status == 200){
          token = response.data.token;
          $cookies.set('token', token);
          window.location.href = '/';          
        }else{
          window.location.href = '/login';          
        }
      })
      .catch(error => {
        window.location.href = '/login';         
        });   
    },
});

appTokenRenew.mount("#root")