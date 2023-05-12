CREATE TABLE produtos(
    id serial primary key,
    nome varchar,
    quantidade int,
    valorUnitario numeric
);

CREATE TABLE usuarios(
    id serial primary key,
    usuario varchar,
    senha varchar
)

INSERT INTO produtos(nome, quantidade, valorUnitario) VALUES
('Sapato', 10, 55.90),
('Blusa azul', 52, 20.00);

INSERT INTO usuarios(usuario, senha) VALUES
('cliente', 'cliente');