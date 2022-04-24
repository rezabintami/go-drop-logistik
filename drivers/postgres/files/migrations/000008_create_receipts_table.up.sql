CREATE TABLE IF NOT EXISTS receipts (
  id serial NOT NULL,
  code varchar(255) NOT NULL,
  receiver varchar(255) NOT NULL,
  phone_receiver varchar(255) NOT NULL,
  address_receiver text NOT NULL,
  sender varchar(255) NOT NULL,
  phone_sender varchar(255) NOT NULL,
  address_sender text NOT NULL,
  weight float NOT NULL,
  unit varchar(100) NOT NULL,
  price float NOT NULL,
  amount float NOT NULL,
  status varchar(255) NOT NULL,
  pickup_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL,
  deleted_at timestamp NULL
);