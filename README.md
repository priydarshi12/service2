Here's an extended API documentation reflecting the `Customer` and `Product` models with your updated structure:

---

## API Documentation

### Customer APIs

#### 1. **Create Customer**
   - **Endpoint:** `POST /customer/create`
   - **Description:** Adds a new customer to the system.
   - **Request Body (JSON):**
     ```json
     {
       "fname": "Priydarshi",
       "lname": "Kumar",
       "email": "kumarpriydarshi6@gmail.com"
     }
     ```
   - **Response:**
     - **Status:** 201 Created
     - **Body:**
       ```json
       {
         "ID": "generated_customer_id",
         "fname": "Priydarshi",
         "lname": "Kumar",
         "email": "kumarpriydarshi6@gmail.com"
       }
       ```
   - **Errors:**
     - 400 Bad Request: Invalid or missing request body fields.
     - 500 Internal Server Error: Could not create customer.

---

#### 2. **Get Customers**
   - **Endpoint:** `GET /customer/get`
   - **Description:** Fetches a list of all customers.
   - **Response:**
     - **Status:** 200 OK
     - **Body:**
       ```json
       [
         {
           "ID": "customer_id_1",
           "fname": "Priydarshi",
           "lname": "Kumar",
           "email": "kumarpriydarshi6@gmail.com"
         },
         {
           "ID": "customer_id_2",
           "fname": "Jane",
           "lname": "Smith",
           "email": "janesmith@example.com"
         }
       ]
       ```
   - **Errors:**
     - 500 Internal Server Error: Unable to fetch customers.

---

#### 3. **Update Customer**
   - **Endpoint:** `PUT /customer/update`
   - **Description:** Updates details of an existing customer.
   - **Request Body (JSON):**
     ```json
     {
       "ID": "customer_id",
       "fname": "Updated First Name",
       "lname": "Updated Last Name",
       "email": "updatedemail@example.com"
     }
     ```
   - **Response:**
     - **Status:** 200 OK
     - **Body:**
       ```json
       {
         "ID": "customer_id",
         "fname": "Updated First Name",
         "lname": "Updated Last Name",
         "email": "updatedemail@example.com"
       }
       ```
   - **Errors:**
     - 400 Bad Request: Invalid request body.
     - 404 Not Found: Customer not found.
     - 500 Internal Server Error: Update operation failed.

---

#### 4. **Delete Customer**
   - **Endpoint:** `DELETE /customer/delete`
   - **Description:** Deletes a customer by ID.
   - **Query Parameter:**
     - `id`: The unique identifier of the customer.
   - **Response:**
     - **Status:** 200 OK
     - **Body:**
       ```json
       {
         "message": "Customer deleted successfully"
       }
       ```
   - **Errors:**
     - 400 Bad Request: Invalid ID format.
     - 404 Not Found: Customer not found.
     - 500 Internal Server Error: Delete operation failed.

---

### Product APIs

#### 1. **Create Product**
   - **Endpoint:** `POST /product/create`
   - **Description:** Adds a new product to the inventory.
   - **Request Body (JSON):**
     ```json
     {
       "name": "Product Name",
       "type": "Product Type",
       "mfd": "2025-01-01",
       "mrp": 100
     }
     ```
   - **Response:**
     - **Status:** 201 Created
     - **Body:**
       ```json
       {
         "ID": "generated_product_id",
         "name": "Product Name",
         "type": "Product Type",
         "mfd": "2025-01-01",
         "mrp": 100
       }
       ```
   - **Errors:**
     - 400 Bad Request: Invalid or missing request body fields.
     - 500 Internal Server Error: Could not create product.

---

#### 2. **Get Products**
   - **Endpoint:** `GET /product/get`
   - **Description:** Fetches a list of all products.
   - **Response:**
     - **Status:** 200 OK
     - **Body:**
       ```json
       [
         {
           "ID": "product_id_1",
           "name": "Product 1",
           "type": "Type 1",
           "mfd": "2024-12-01",
           "mrp": 100
         },
         {
           "ID": "product_id_2",
           "name": "Product 2",
           "type": "Type 2",
           "mfd": "2024-11-01",
           "mrp": 200
         }
       ]
       ```
   - **Errors:**
     - 500 Internal Server Error: Unable to fetch products.

---

#### 3. **Update Product**
   - **Endpoint:** `PUT /product/update`
   - **Description:** Updates details of an existing product.
   - **Request Body (JSON):**
     ```json
     {
       "ID": "product_id",
       "name": "Updated Product Name",
       "type": "Updated Product Type",
       "mfd": "2025-01-01",
       "mrp": 150
     }
     ```
   - **Response:**
     - **Status:** 200 OK
     - **Body:**
       ```json
       {
         "ID": "product_id",
         "name": "Updated Product Name",
         "type": "Updated Product Type",
         "mfd": "2025-01-01",
         "mrp": 150
       }
       ```
   - **Errors:**
     - 400 Bad Request: Invalid request body.
     - 404 Not Found: Product not found.
     - 500 Internal Server Error: Update operation failed.

---

#### 4. **Delete Product**
   - **Endpoint:** `DELETE /product/delete`
   - **Description:** Deletes a product by ID.
   - **Query Parameter:**
     - `id`: The unique identifier of the product.
   - **Response:**
     - **Status:** 200 OK
     - **Body:**
       ```json
       {
         "message": "Product deleted successfully"
       }
       ```
   - **Errors:**
     - 400 Bad Request: Invalid ID format.
     - 404 Not Found: Product not found.
     - 500 Internal Server Error: Delete operation failed.

---

