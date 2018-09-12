CREATE TABLE mutant.mutants (
  id integer unique NOT NULL AUTO_INCREMENT,
  dna varchar(200) not null,
  is_mutant BOOLEAN default 0
);