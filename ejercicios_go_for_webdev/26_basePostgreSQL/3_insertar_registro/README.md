# insertar un registro
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (1, 'Mark', 7, '1212 E. Lane, Someville, AK, 57483', 43000.00 ,'1992-01-13');
```

## listar registros en una tabla
```
SELECT * FROM <table name>;
```

## insertar un registro - variaciones
los valores omitidos tendrán el [valor predeterminado](https://www.postgresql.org/docs/9.3/static/ddl-default.html):
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,BDAY) VALUES (2, 'Marian', 8, '7214 Wonderlust Ave, Lost Lake, KS, 22897', '1989-11-21');
```

podemos usar DEFAULT en lugar de dejar un campo en blanco o especificar un valor:
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (3, 'Maxwell', 6, '7215 Jasmine Place, Corinda, CA 98743', 87500.00, DEFAULT);
```

podemos insertar múltiples filas:
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (4, 'Jasmine', 5, '983 Star Ave., Brooklyn, NY, 00912 ', 55700.00, '1997-12-13' ), (5, 'Orranda', 9, '745 Hammer Lane, Hammerfield, Texas, 75839', 65350.00 , '1992-12-13');
```
