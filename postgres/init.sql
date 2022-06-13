DROP TABLE IF EXISTS bagian;
CREATE SEQUENCE bagian_id START 1;
CREATE TABLE bagian (
    ID serial PRIMARY KEY,
    KODEBAGIAN TEXT,
    NAMABAGIAN TEXT,
    KETERANGAN TEXT,
);

DROP TABLE IF EXISTS subbag;
CREATE SEQUENCE subbag_id START 1;
CREATE TABLE subbag (
    ID serial PRIMARY KEY,
    KODEBAGIAN TEXT,
    KDSUBBAG TEXT,
    NAMABAGIAN TEXT,
    KETERANGAN TEXT,
);

DROP TABLE IF EXISTS tiketantrian;
CREATE SEQUENCE tiketantrian_id START 1;
CREATE TABLE tiketantrian (
    ID serial PRIMARY KEY,
    TGLBUATTIKET TEXT,
    STATUSTIKET TEXT,
    STATUSCETAK TEXT,
);


DROP TABLE IF EXISTS tikettamu;
CREATE SEQUENCE tikettamu_id START 1;
CREATE TABLE tikettamu (
    ID serial PRIMARY KEY,
    NOTIKET TEXT,
    NOIDENTITAS TEXT,
    NAMA TEXT,
    KDBAGSUBSEKSI TEXT,
    KDSUBBAGSEKSI TEXT,
    JABATAN TEXT,
    KEPERLUANBERTAMU TEXT,
    TGLMINTABERTAMU TEXT,
    MINTAJAMBERTAMU TEXT,
);