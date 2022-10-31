var Vue = new Vue({
    el: '#root',
    delimiters: ['@{', '}'],
    data: {
    supplier: {id: 0, name: '', phone: '', address: '', notes: ''},
    suppliers: [],
    isUpdate: false,
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
          this.$http.post('supplier/create',{
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
        this.$http.put('supplier/' + supplier.id,{
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
      this.$http.get('supplier/list').then(response =>{
        if(response.status == 200){
          this.suppliers =  response.body;
          this.supplier = {id: 0, name: '', phone: '', address: '', notes: ''};
        }
      });
    },
    getDetailSupplier(supplier, supplierIndex){
      this.isUpdate = true;
      this.changeHeader();
      this.$http.get('supplier/'+supplier.id).then(response =>{
        if(response.status == 200){
          this.supplier =  response.body;
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
        this.$http.delete('supplier/'+ supplier.id).then(response =>{
          if(response.status == 200){
            this.suppliers.splice(supplierIndex,1);
            this.supplier = {id: 0, name: '', phone: '', address: '', notes: ''};
          }
        });
      }
    },
  },
  mounted() { 
    this.listSupplier();
  }
});