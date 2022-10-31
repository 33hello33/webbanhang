var Vue = new Vue({
    el: '#root',
    delimiters: ['@{', '}'],
    data: {
    customer: {id: 0, name: '', phone: '', address: '', notes: ''},
    customers: [],
    isUpdate: false,
  },
  methods: {
    addCustomer() {
      this.isUpdate = false;
      this.changeHeader();
      this.customer = {id: 0, name: '', phone: '', address: ''};
    },
    createCustomer(customer, customerIndex) {
      if(this.isUpdate == false){
          // create new customer
          this.$http.post('customer/create',{
            name: customer.name, 
            address: customer.address, 
            phone: customer.phone}).then(response => {
            if(response.status == 200){
              this.listCustomer();
            }
          });
      }else{
        // update customer
        this.$http.put('customer/' + customer.id,{
            id: customer.id,
            name: customer.name, 
            address: customer.address, 
            phone: customer.phone}).then(response => {
            if(response.status == 200){
              this.listCustomer();
            }
        });
      }
    },
    checkForEnter(event){
        if (event.key == "Enter") {
          this.createCustomer();
        }
    },
    listCustomer(){
      this.$http.get('customer/list').then(response =>{
        if(response.status == 200){
          this.customers =  response.body;
          this.customer = {id: 0, name: '', phone: '', address: '', notes: ''};
        }
      });
    },
    getDetailCustomer(customer, customerIndex){
      this.isUpdate = true;
      this.changeHeader();
      this.$http.get('customer/'+customer.phone).then(response =>{
        if(response.status == 200){
          this.customer =  response.body;
        }
      })
    },
    changeHeader(){
      if(this.isUpdate == true){
        this.$refs.headerCustomer.innerText = "Sửa Khách Hàng";
        this.$refs.buttonCustomer.innerText = "Sửa Khách Hàng";
      }else{
        this.$refs.headerCustomer.innerText = "Thêm Khách Hàng";
        this.$refs.buttonCustomer.innerText = "Thêm Khách Hàng";
      }
    },
    deleteCustomer(customer, customerIndex){
      if(confirm("Are you sure ?")){
        this.$http.delete('customer/'+ customer.id).then(response =>{
          if(response.status == 200){
            this.customers.splice(customerIndex,1);
            this.customer = {id: 0, name: '', phone: '', address: '', notes: ''};
          }
        });
      }
    },
  },
  mounted() { 
    this.listCustomer();
  }
});