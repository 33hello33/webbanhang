var Vue = new Vue({
    el: '#root',
    delimiters: ['@{', '}'],
    data: {
      //invo: {id: 0, created_at: '', total_money: 0, had_paid: 0, cutomers_name: '', customers_phone:''},
      invoices: null,
    },
    mounted(){
      this.listInvoices();
    },
    methods: {
      listInvoices(){
        this.$http.get('invoice/list').then(response =>{
          if(response.status == 200){
            this.invoices = response.body;
          }else{
            console.log(response.body);
          }
        });
      },
      getDetailInvoice(invoiceIndex){
        this.$http.get('invoice/' + invoice.id).then(response =>{
          if (response.status == 200){

          }
        })
      },
    }});