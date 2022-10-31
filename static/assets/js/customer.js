const { createApp } = Vue;
const appCustomer = createApp({
    delimiters: ['@{', '}'],
    data(){
      return {
        customer: {id: 0, name: '', phone: '', address: {String: '', Valid: false}},
        customers: [],
        isUpdate: false,
      }
  },
  methods: {
    addCustomer() {
      this.isUpdate = false;
      this.changeHeader();
      this.customer = {id: 0, name: '', phone: '', address: {String: '', Valid: false}};
    },
    createCustomer(customer, customerIndex) {
      if(this.isUpdate == false){
          // create new customer
          axios.post('customer/create',{
            name: customer.name, 
            address: customer.address.String, 
            phone: customer.phone}).then(response => {
            if(response.status == 200){
              this.listCustomer();
            }
          });
      }else{
        // update customer
        axios.put('customer/' + customer.id,{
            id: customer.id,
            name: customer.name, 
            address: customer.address.String, 
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
      axios.get('customer/list').then(response =>{
        if(response.status == 200){
          this.customers =  response.data;
          this.customer = {id: 0, name: '', phone: '', address: {String: '', Valid: false}};
        }
      });
    },
    getDetailCustomer(customer, customerIndex){
      this.isUpdate = true;
      this.changeHeader();
      axios.get('customer/'+customer.id).then(response =>{
        if(response.status == 200){
          this.customer =  response.data;
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
        axios.delete('customer/'+ customer.id).then(response =>{
          if(response.status == 200){
            this.customers.splice(customerIndex,1);
            this.customer = {id: 0, name: '', phone: '', address: {String: '', Valid: false}};
          }
        });
      }
    },
  },
  beforeMount(){
    this.listCustomer();
 },
})  
appCustomer.mount("#root")