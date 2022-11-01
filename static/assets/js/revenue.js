appRevenue = createApp({
    delimiters: ['@{', '}'],
    data(){
      return {
        //invo: {id: 0, created_at: '', total_money: 0, had_paid: 0, cutomers_name: '', customers_phone:''},
        invoices: null,
      }
    },
    methods: {
      listInvoices(){
        axios.get('invoice/list').then(response =>{
          if(response.status == 200){
            this.invoices = response.data;
          }
        })
        .catch(error => {
          alert(error.data.Error);
          });   
      },
      getDetailInvoice(invoiceIndex){
        axios.get('invoice/' + invoice.id).then(response =>{
          if (response.status == 200){

          }
        })
      },
    },
    beforeMount(){
      this.listInvoices();
    },
  });
  appRevenue.mount("#root")