var Vue = new Vue({
  el: '#root1',
  delimiters: ['@{', '}'],
  data: {
    //invo: {id: 0, created_at: '', total_money: 0, had_paid: 0, cutomers_name: '', customers_phone:''},
    invoices: null,
  },
  methods: {
    deleteToken(){
      var refresh_token = localStorage.getItem('refresh_token');
    
      this.$http.post('/token/delete', {"refresh_token" : refresh_token}).then( response =>{
        if(response.status == 200){
          this.invoices = response.body;
        }
      });
    },
  }});