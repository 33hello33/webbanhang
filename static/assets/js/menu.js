const {createApp} = Vue 
appMenu = createApp({
  methods:{
    logoutUser(){
      axios.post('logout', {'refresh_token': localStorage.getItem('refresh_token')}).then(response => {
        if(response.status == 200)
          console.log('logout successfully');
        });      
    },
    changeColorMenu(id){
      var modal = document.getElementById(id);
      modal.style.backgroundColor = "#04AA6D";
    },
  },
  mounted(){
    var url = window.location.href;
    if(url.includes("invoice")){
      this.changeColorMenu("invoice");
    }else if(url.includes("product")){
      this.changeColorMenu("product");
    }else if(url.includes("supplier")){
      this.changeColorMenu("supplier");
    }else if(url.includes("customer")){
      this.changeColorMenu("customer");
    }else if(url.includes("revenue")){
      this.changeColorMenu("revenue");
    }else if(url.includes("statistic")){
      this.changeColorMenu("statistic");
    }
  },
})
appMenu.mount('#topnav');