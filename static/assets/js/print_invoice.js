const {createApp} = Vue 
appRevenue = createApp({
    delimiters: ['@{', '}'],
    data(){
      return {
        invoice: {id: 0, created_at: '', total_money: 0, had_paid: 0, name: '', phone:'', is_done:''},
        productTbls: [],
      }
    },
    methods: {
    getInvoice(id){
        axios.get('/invoice/' + id)
        .then(response => {
            console.log(response.data);
            this.invoice = response.data;
        })
        .catch(error => {
            console.log(error.data.Error);
        })
    },
    getDetailInvoice(id){
        axios.get('/invoice/detail/' + id)
        .then(response => {
          if(response.status==200){
            this.productTbls = response.data;
          }
        })
        .catch(error => {
            console.log(error.data.Error);
          });   
      },
    },
    beforeMount(){
        const url = window.location.href;
        const id = url.split("/").slice(-1)[0];
        this.getInvoice(id);
        this.getDetailInvoice(id);
    },
  });
  appRevenue.mount("#root")