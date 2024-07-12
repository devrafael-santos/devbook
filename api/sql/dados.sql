insert into usuarios (nome, nick, email, senha) 
value
("Rafael", "raffas", "rafael@gmail.com", "$2a$10$K.muryjC.PnQEG/XrayYiuEVSoKEZp1ESaG60OHdKWu5zzQjefZIK"),
("Paulo", "paulooo", "paulo@Tgmail.com", "$2a$10$K.muryjC.PnQEG/XrayYiuEVSoKEZp1ESaG60OHdKWu5zzQjefZIK"),
("Rodrigo", "roggg", "rodr@gmail.com", "$2a$10$K.muryjC.PnQEG/XrayYiuEVSoKEZp1ESaG60OHdKWu5zzQjefZIK");

insert into seguidores (usuario_id, seguidor_id)
value
(1, 2),
(1, 3),
(3, 1);