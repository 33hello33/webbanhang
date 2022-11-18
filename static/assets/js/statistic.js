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
        selling_product: {name:"", selling_amount:""},
        top_selling_products: [],
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
            this.total_price_all_invoice = response.data.sum_total;
            console.log('total money');
            console.log(response.data.sum_total);
          }
        })
        .catch(error => {
          console.log(error.data.Error);
          });    
      },
      getTopSellingProduct(number, from_date, to_date){
        axios.post('statistic/top_selling_product', {
          'number_top':number,
          'from_date':  String(from_date), 
          'to_date':  String(to_date),
        })
        .then(response =>{
          var top_products = response.data;
          this.top_selling_products = response.data;
        })
        .catch(error => {
          console.log(error.data);
          });   
      },
      findAll(){
        this.findInvoices(this.from_date, this.to_date, 'Tất cả', '', 'Tất cả');
        this.getTopSellingProduct(5, this.from_date, this.to_date);
      },
    },
    beforeMount(){
      this.from_date = this.formatDate();
      this.to_date = this.formatDate();
      this.findInvoices(this.from_date, this.to_date, 'Tất cả', '', 'Tất cả');
      this.getTopSellingProduct(5, this.from_date, this.to_date);
      this.filter_by_id = this.filter_ids[0];
      this.filter_by_status = this.filter_status[0];
    },
  });
  appRevenue.mount("#root")