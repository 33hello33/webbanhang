appRevenue = createApp({
    delimiters: ['@{', '}'],
    data(){
      return {
        from_date: '',
        to_date: '',
        invoice: {id: 0, created_at: '', total_money: 0, had_paid: 0, cutomer_name: '', customer_phone:''},
        invoices: [],
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
          console.log(error.data.Error);
          });   
      },
      getDetailInvoice(invoiceIndex){
        axios.get('invoice/' + invoice.id).then(response =>{
          if (response.status == 200){

          }
        })
      },
      formatDate() {
        var d = new Date(),
            month = '' + (d.getMonth() + 1),
            day = '' + d.getDate(),
            year = d.getFullYear();
    
        if (month.length < 2) 
            month = '0' + month;
        if (day.length < 2) 
            day = '0' + day;
    
        return [year, month, day].join('-');
      },
      findInvoices(){
        axios.post('invoice/find', {'from_date': String(this.from_date), 'to_date': String(this.to_date)})
        .then(response => {
          if(response.status == 200){
            this.invoices = response.data;
          }
        })
        .catch(error => {
          console.log(error.data.Error);
          });   
      }
    },
    beforeMount(){
      this.listInvoices();
      this.from_date = this.formatDate();
      this.to_date = this.formatDate();
    },
    
  });
  appRevenue.mount("#root")