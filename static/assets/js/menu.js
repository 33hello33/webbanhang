const {createApp} = Vue 
appMenu = createApp({
  methods:{
  logoutUser(){
    axios.post('logout', {'refresh_token': localStorage.getItem('refresh_token')}).then(response => {
      if(response.status == 200)
        console.log('logout successfully');
      });      
  },
  }
})
appMenu.mount('#topnav')