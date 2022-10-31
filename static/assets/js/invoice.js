var Vue = new Vue({
    el: '#root',
    delimiters: ['@{', '}'],
    data: {
      customer: {name: '', phone: '', address:''},
      invoice:{customer_phone: '', total_money: 0, had_paid:0},
      product: {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0},
      products: [],
      productTbl: {id: 0, name: '', unit: '', amount: 0, price: 0, total_price: 0, discount: 0, last_price: 0},
      productTbls: [],
      resetKey: 0,
      total_money_to_pay: 0,
    },
    mounted(){
      this.listProduct();
    },
    methods: {
      listProduct(){
        this.$http.get('product/list').then(response =>{
          if(response.status == 200){
            this.products = response.body;
            this.product = {id: 0, name: '', unit: '', amount: 0, price: 0, price_import: 0, warehouse: '', id_supplier: 0};
          }else{
            console.log(response.body);
          }
        });
      },
      getDetailProduct(product, productIndex){
        this.$http.get('product/' + product.id).then(response =>{
          if (response.status == 200){
            this.productTbl = response.body;
            this.productTbl.amount = 1;
            this.productTbl.discount = 0;
            this.productTbl.total_price = this.productTbl.price;
            this.productTbl.last_price = this.productTbl.price;
            this.productTbls.push(this.productTbl);
            this.calcToTalPriceToPay();
          }
        })
      },
      onchangeTable(index) {
        this.productTbls[index].total_price = this.productTbls[index].price * this.productTbls[index].amount;
        this.productTbls[index].last_price = this.productTbls[index].total_price * ( 1- this.productTbls[index].discount/100);
        this.calcToTalPriceToPay();
        this.resetKey +=1;  
      },
      deleteRow(productindexTbl){
        this.productTbls.splice(productindexTbl,1);
        this.calcToTalPriceToPay();
      },
      calcToTalPriceToPay(){
        sum = 0;
        for(pd of this.productTbls){
          sum+= pd.last_price;
        }
        this.total_money_to_pay = sum;
      },
      paidInvoice(){
        // send info customer, if not exist, create new customer
        this.invoice.customer_phone = this.customer.phone;
        this.invoice.total_money = this.total_money_to_pay;
        this.$http.post('customer', this.customer).then(response =>{
          if(response.status = 200){
            // send info invoice, all product of invoice
            this.$http.post('invoice', {"invoice" : this.invoice, "products": this.productTbls}).then(response =>{
              if(response.status = 200){
                this.reset();
              }else{
                console.log(response.body);
              }
            })
          }else{
            console.log(response.body);
          }
        })
      },
      reset(){
        this.productTbls = [];
        this.total_money_to_pay = 0;
      },
    }});