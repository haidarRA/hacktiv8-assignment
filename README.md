# hacktiv8-assignment

# 🛍️ Assignment: Membangun REST API E-Commerce (Produk, Sumber, dan Transaksi)

## Product

Daftar endpoint dan method:
1. `GET /products` – Ambil seluruh produk
   <img width="1354" height="718" alt="image" src="https://github.com/user-attachments/assets/69c424b9-608d-4c8e-a087-5bca2dff4c3f" />   

2. `GET /products/:id` – Ambil produk berdasarkan ID
   <img width="1367" height="569" alt="image" src="https://github.com/user-attachments/assets/ff6885cb-7b0f-45a9-9aed-af018f9af73b" />

   Jika produk tidak ada:
   <img width="1362" height="332" alt="image" src="https://github.com/user-attachments/assets/658531dc-190d-4562-aac5-0a202530a1f8" />
  
3. `POST /products` – Tambah produk
   <img width="1381" height="807" alt="image" src="https://github.com/user-attachments/assets/d503e55c-f92d-48c5-92b0-7d9d704800f4" />

   Jika nama produk kosong dan/atau harga dan stock tidak valid:
   <img width="1389" height="650" alt="image" src="https://github.com/user-attachments/assets/98190f02-4a0e-40fa-a4e0-2b887222f350" />

4. `PUT /products/:id` – Ubah produk
   <img width="1359" height="645" alt="image" src="https://github.com/user-attachments/assets/56a0ffab-ebdb-4ba1-abf3-5963d2bb3e33" />

   Jika nama produk kosong dan/atau harga dan stock tidak valid:
   <img width="1369" height="486" alt="image" src="https://github.com/user-attachments/assets/0852e013-95a6-4746-b074-a84eff0cfc47" />

   Jika produk tidak ada:
   <img width="1359" height="325" alt="image" src="https://github.com/user-attachments/assets/3eddf33f-810f-4127-8c67-5835f13fb91c" />

5. `DELETE /products/:id` – Hapus produk
   <img width="1361" height="372" alt="image" src="https://github.com/user-attachments/assets/c489ed78-a816-422d-8546-870e139bca30" />

   Jika produk tidak ada:
   <img width="1363" height="398" alt="image" src="https://github.com/user-attachments/assets/9270fb7b-5317-45ba-9a7b-2dd1a062bae0" />

## Source

1. `GET /sources` – Ambil seluruh source
   <img width="1367" height="523" alt="image" src="https://github.com/user-attachments/assets/12d5fd16-99d5-4839-93d4-2b11bbbbcb41" />

2. `GET /sources/:id` – Ambil detail source
   <img width="1375" height="486" alt="image" src="https://github.com/user-attachments/assets/aeba5047-0cec-4947-8bc5-9f15c7f45aa5" />

   Jika produk tidak ada:
   <img width="1368" height="320" alt="image" src="https://github.com/user-attachments/assets/00da89b4-a774-47e2-950b-0d9551d7315b" />

3. `POST /sources` – Tambah source
   <img width="1368" height="563" alt="image" src="https://github.com/user-attachments/assets/c9f65264-5d44-43cd-900b-b0aa80d90cd0" />

4. `PUT /sources/:id` – Ubah source
   <img width="1366" height="571" alt="image" src="https://github.com/user-attachments/assets/a701461a-aa6f-477f-9ddd-a59c7d9736c2" />

   Jika source tidak ada:
   <img width="1369" height="324" alt="image" src="https://github.com/user-attachments/assets/95ac812f-0dcd-4054-b824-59beaf3dbdde" />

5. `DELETE /sources/:id` – Hapus source
   <img width="1375" height="334" alt="image" src="https://github.com/user-attachments/assets/e13ce7f5-a4e3-4401-8aee-22ecc546975b" />

   Jika source tidak ada:
   <img width="1368" height="323" alt="image" src="https://github.com/user-attachments/assets/285c89dd-aaa5-4968-a9c7-ff00a5630948" />

## Transaction

1. `POST /transactions` – Tambah transaksi (mengurangi stok produk)
   <img width="1372" height="591" alt="image" src="https://github.com/user-attachments/assets/7044f233-c8c3-49a4-ada8-9d0991eb14d6" />

   Setelah transaction:
   <img width="1363" height="533" alt="image" src="https://github.com/user-attachments/assets/10dd1bc2-9b6a-4fdb-94ad-1527cee1c08c" />

   Jika produk tidak ada:
   <img width="1365" height="483" alt="image" src="https://github.com/user-attachments/assets/0eeb41da-7644-453e-b4ea-96f223c178d0" />

   Jika source produk tidak valid:
   <img width="1363" height="712" alt="image" src="https://github.com/user-attachments/assets/2bf8c5fb-2f50-4e9a-8b9d-58645a3953f2" />
   <img width="1359" height="439" alt="image" src="https://github.com/user-attachments/assets/880737d3-0919-44e8-acd9-b49d79abf639" />
   <img width="1360" height="484" alt="image" src="https://github.com/user-attachments/assets/4c2b6a5c-5801-4df2-a2bf-e0973191e1ed" />

   Jika quantity melebihi stock:
   <img width="1372" height="484" alt="image" src="https://github.com/user-attachments/assets/c4c9760f-1c61-4394-a72e-723abed03fcf" />

2. `GET /transactions` – Ambil daftar transaksi
   <img width="1381" height="653" alt="image" src="https://github.com/user-attachments/assets/105c914b-bbcc-41ed-9be5-3015fc9ce283" />

3. `GET /transactions/:id` – Ambil detail transaksi
   <img width="1369" height="448" alt="image" src="https://github.com/user-attachments/assets/ad2bc6e0-5f9b-4eba-8880-905b19984c89" />

   Jika transaksi tidak ada:
   <img width="1360" height="322" alt="image" src="https://github.com/user-attachments/assets/042f7691-a3d9-404c-81c5-fca544c4712b" />
