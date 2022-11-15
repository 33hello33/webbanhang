appRevenue = createApp({
    delimiters: ['@{', '}'],
    data(){
      return {
        from_date: '',
        to_date: '',
        filter_by_id: '',
        filter_ids: ['Tất cả','Mã đơn hàng','Tên khách hàng'],
        filter_input: '',
        filter_by_status: '',
        filter_status: ['Tất cả', 'Hoàn thành', 'Nợ'],
        total_price_all_invoice: 0,
        invoice: {id: 0, created_at: '', total_money: 0, had_paid: 0, name: '', phone:'', is_done:''},
        invoices: [],
        productTbls: [],
      }
    },
    methods: {
      onChangeFilterID(){
        if(this.filter_by_id == this.filter_ids[0]){
          this.$refs.filter_input.placeholder = '';
        }else if (this.filter_by_id == this.filter_ids[1]){
          this.$refs.filter_input.placeholder = 'Nhập mã đơn hàng';
        }else{
          this.$refs.filter_input.placeholder = 'Nhập tên khách hàng';
        }
      },
      onChangeFilterStatus(){
        console.log(this.filter_by_status);
      },
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
      findInvoices(from_date, to_date, filter_by_id, filter_input, filter_by_status){
        axios.post('invoice/find', {
          'from_date': String(from_date),
           'to_date': String(to_date),
           'filter_by_id': filter_by_id,
           'filter_input': filter_input,
           'filter_by_status': filter_by_status,
          })
        .then(response => {
          if(response.status == 200){
            this.invoices = response.data.invoices;
            this.total_price_all_invoice = response.data.sum_total;

            this.invoices.forEach(invoice => {
              if(invoice.is_done == true){
                invoice.is_done = 'Hoàn thành';
              }else{
                invoice.is_done = 'Nợ';
              }
            });
          }
        })
        .catch(error => {
          console.log(error.data.Error);
          });    
      },
      showDetailInvoice(invoice){
        this.invoice = invoice;
        var modal = document.getElementById("DetailInvoice");
        modal.style.display = "block";
        axios.get('invoice/'+invoice.id)
        .then(response => {
          if(response.status==200){
            this.productTbls = response.data;
          }
        })
        .catch(error => {
          console.log(error.data.Error);
          });   
      },
      closeDetailInvoice(){
        var modal = document.getElementById("DetailInvoice");
        modal.style.display = "none";
      },
    },
    beforeMount(){
      this.from_date = this.formatDate();
      this.to_date = this.formatDate();
      this.findInvoices(this.from_date, this.to_date, "", "", "");
      this.filter_by_id = this.filter_ids[0];
      this.filter_by_status = this.filter_status[0];
    },
  });
  appRevenue.mount("#root")