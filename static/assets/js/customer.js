const appCustomer = createApp({
    delimiters: ['@{', '}'],
    data(){
      return {
        customer: {id: 0, name: '', phone: '', address: {String: '', Valid: false}},
        customers: [],
        isUpdate: false,
        searchInput:'',
        timeOut: 0,
      }
  },
  methods: {
    searchCustomer(){
      if(this.searchInput == ""){
        this.listCustomer();
        return;
      }

      axios.get('customer/search/'+this.searchInput)
      .then(response=>{
        if(response.status == 200){
          this.customers =  response.data;
        }
      })
      .catch(error => {
        console.log('get list customer err: ' + error.data.Error);
      });   
    },
    searchInputChanges(){
      clearTimeout(this.timeOut);

      this.timeOut = setTimeout(() => {
        this.searchCustomer();
      }, 300);
    },

    addCustomer() {
      this.isUpdate = false;
      this.changeHeader();
      this.customer = {id: 0, name: '', phone: '', address: {String: '', Valid: false}};
    },
    createCustomer(customer, customerIndex) {
      if(this.isUpdate == false){
          // create new customer
          if (customer.address.String) {
            customer.address.Valid = true;
          }
          axios.post('customer/create', customer)
          .then(response => {
            if(response.status == 200){
              this.listCustomer();
            }
          });
      }else{
        // update customer
        if (customer.address.String) {
          customer.address.Valid = true;
        }
        axios.put('customer/' + customer.id, customer)
        .then(response => {
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
          this.customer = response.data;
        }
      })
    },
    changeHeader(){
      if(this.isUpdate == true){
        this.$refs.headerCustomer.innerText = "S???a Kh??ch H??ng";
        this.$refs.buttonCustomer.innerText = "S???a Kh??ch H??ng";
      }else{
        this.$refs.headerCustomer.innerText = "Th??m Kh??ch H??ng";
        this.$refs.buttonCustomer.innerText = "Th??m Kh??ch H??ng";
      }
    },
    deleteCustomer(customer, customerIndex){
      if(confirm("Are you sure ?")){
        axios.delete('customer/'+ customer.id)
        .then(response =>{
          if(response.status == 200){
            this.customers.splice(customerIndex,1);
            this.customer = {id: 0, name: '', phone: '', address: {String: '', Valid: false}};
          }
        })
        .catch(error=>{
          if(error.data.Error.includes('foreign key')){
            alert("Kh??ng th??? x??a kh??ch h??ng v?? kh??ch h??ng n??y ??ang t???n t???i trong h??a ????n");
          }else{
            console.log(error.data.Error);
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