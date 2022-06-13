INSER INTO tikettamu (notiket, noidentitas, nama, kdbagsubseksi, kdsubbagseksi, jabatan, keperluanbertamu, tglmintabertamu, mintajambertamu)
VALUE ('D01-02-01-001-10-05-21', '3273221306720011', 'Trit-Tux', 'D01', 'D01-01', 'Kepala Sub Bagian Tata Usaha', '', '12902389283', '32003284');

SELECT * FROM bagian;
SELECT * FROM subbag;
SELECT * FROM tiketantrian;
SELECT * FROM tikettamu;

SELECT b.*, s.* FROM bagian b, 
JOIN subbag s ON s.kodebagian = b.kodebagian

SELECT * FROM tikettamu t
JOIN bagian b ON b.kodebagian = t.kdbagsubseksi
JOIN subbag s ON s.kdsubbag = t.kdsubbagseksi

DELETE FROM tikettamu WHERE id=1

SELECT * FROM tikettamu WHERE kdbagsubseksi = 'D01'

