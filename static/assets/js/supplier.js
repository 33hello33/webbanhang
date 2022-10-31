const {createApp} = Vue
appSupplier = createApp({
    delimiters: ['@{', '}'],
    data() {
      return{
        supplier: {id: 0, name: '', phone: '', address: '', notes: ''},
        suppliers: [],
        isUpdate: false,
      }
  },
  methods: {
    addSupplier() {
      this.isUpdate = false;
      this.changeHeader();
      this.supplier = {id: 0, name: '', phone: '', address: '', notes: ''};
    },
    createSupplier(supplier, supplierIndex) {
      if(this.isUpdate == false){
          // create new supplier
          axios.post('supplier/create',{
            name: supplier.name, 
            address: supplier.address, 
            phone: supplier.phone, 
            notes: supplier.notes}).then(response => {
            if(response.status == 200){
              this.listSupplier();
            }
          });
      }else{
        // update supplier
        axios.put('supplier/' + supplier.id,{
            id: supplier.id,
            name: supplier.name, 
            address: supplier.address, 
            phone: supplier.phone, 
            notes: supplier.notes}).then(response => {
            if(response.status == 200){
              this.listSupplier();
            }
        });
      }
    },
    checkForEnter(event){
        if (event.key == "Enter") {
          this.createSupplier();
        }
    },
    listSupplier(){
      axios.get('supplier/list').then(response =>{
        if(response.status == 200){
          this.suppliers =  response.data;
          this.supplier = {id: 0, name: '', phone: '', address: '', notes: ''};
        }
      });
    },
    getDetailSupplier(supplier, supplierIndex){
      this.isUpdate = true;
      this.changeHeader();
      axios.get('supplier/'+supplier.id).then(response =>{
        if(response.status == 200){
          this.supplier =  response.data;
        }
      })
    },
    changeHeader(){
      if(this.isUpdate == true){
        this.$refs.headerSupplier.innerText = "Sửa Nhà Cung Cấp";
        this.$refs.buttonSupplier.innerText = "Sửa Nhà Cung Cấp";
      }else{
        this.$refs.headerSupplier.innerText = "Thêm Nhà Cung Cấp";
        this.$refs.buttonSupplier.innerText = "Thêm Nhà Cung Cấp";
      }
    },
    deleteSupplier(supplier, supplierIndex){
      if(confirm("Are you sure ?")){
        axios.delete('supplier/'+ supplier.id).then(response =>{
          if(response.status == 200){
            this.suppliers.splice(supplierIndex,1);
            this.supplier = {id: 0, name: '', phone: '', address: '', notes: ''};
          }
        });
      }
    },
  },
  beforeMount() { 
    this.listSupplier();
  }
});