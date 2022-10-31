const {createApp} = Vue
appTokenRenew = createApp({
    delimiters: ['@{', '}'],
    beforeMount() { 
      var refresh_token = localStorage.getItem('refresh_token');
      
      axios.post('/token/renew', {"refresh_token" : refresh_token}).then( response =>{
        if(response.status == 200){
          token = response.data.token;
          document.cookie ='token= '+ token;
          window.location.href = '/';          
        }
      })
    },
});

appTokenRenew.mount("#root")