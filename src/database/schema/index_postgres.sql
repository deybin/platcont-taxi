
CREATE TABLE Taxi_ClientesCars (
	c_plac varchar(7) NOT NULL PRIMARY KEY,
	n_docu varchar(11) NOT NULL,
	l_marc varchar(50) DEFAULT '' ,
	l_mode varchar(50) DEFAULT '' ,
	l_color varchar(70) DEFAULT '' ,
	n_year int NOT NULL ,
	c_mode varchar(4) DEFAULT '' ,
	n_seri varchar(17) DEFAULT '',
	n_pasa int DEFAULT '' ,
	l_obse varchar(100) DEFAULT '' ,
	k_stad int DEFAULT 0 ,
	FOREIGN KEY (n_docu) REFERENCES Requ_Clientes(n_docu)
);

CREATE TABLE Taxi_Servicios (
	id_serv varchar(36)  NOT NULL PRIMARY KEY,
	n_year int  NOT NULL,
	n_month int  NOT NULL,
	n_docu varchar(11)  DEFAULT '',
	f_fact varchar(10)  DEFAULT '',
	s_impo float8 DEFAULT 0,
	c_plac varchar(6)  DEFAULT '',
	k_stad int DEFAULT 0,
	f_digi timestamp DEFAULT now(),
	CONSTRAINT Taxi_Servicios_n_docu_c_plac UNIQUE (n_docu,c_plac),
	FOREIGN KEY (c_plac) REFERENCES ClientesCars(c_plac),
	FOREIGN KEY (n_docu) REFERENCES Clientes(n_docu)
);

CREATE TABLE Taxi_ServiciosDetalle (
	id_serv varchar(36)  NOT NULL,
	n_year int  NOT NULL,
	n_month int  NOT NULL,
	f_pago varchar(10) DEFAULT '',
	s_impo float8 DEFAULT 0,
	k_stad int DEFAULT 0,
	FOREIGN KEY (id_serv) REFERENCES Servicios(id_serv)
	CONSTRAINT Taxi_ServiciosDetalle_c_year_c_month UNIQUE (n_year,n_month)
);

