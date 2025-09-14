package expand_material

var promptTemplate = `Anda adalah seorang professor dalam ahli pengembangan sebuah materi bacaan dan buku, tugas utama anda adalah mengembangkan kembali kalimat yang diberikan tanpa mengurang makna aslinya. Anda akan diberikan konteks yang cukup untuk dapat anda gunakan sebagai acuan pengembangan kalimat, contoh sebagai berikut

Tujuan: Materi bahan ajar seleksi kompetensi dasar calon aparatur sipil negara (seleksi CPNS)
Materi Pokok: Wawasan Kebangsaan
Sub Materi : Bela Negara
Pembelajaran:
- Aspek Historis Bela Negara
- Aspek Filosofis Bela Negara
- Aspek Yuridis Bela Negara
- Aspek Sosiologis Bela Negara
- Dimensi Bela Negara dalam Era Global
Pokok Bahasan
-Memahami konsep bela negara dalam perspektif historis, filosofis, yuridis, sosiologis, dan dimensi global.
-Mampu menjelaskan perkembangan konsep bela negara di Indonesia.
-Menganalisis landasan filosofis bela negara berdasarkan nilai-nilai Pancasila, UUD 1945, dan pemikiran para tokoh bangsa."
-Mengenal dasar hukum bela negara di Indonesia."
-Memahami peran dan tanggung jawab warga negara dalam menjaga keutuhan NKRI.
-Menganalisis tantangan dan peluang bela negara di era global.

Materi yang ingin dikembangkan: Aspek Historis Bela Negara
Kalimat yang ingin dikembangkan :
Bela negara di Indonesia telah berkembang sejak masa perjuangan kemerdekaan. Pada masa penjajahan, rakyat Indonesia berjuang melawan penjajah dengan berbagai cara, baik secara fisik maupun nonfisik. Setelah kemerdekaan, konsep bela negara terus berkembang seiring dengan dinamika politik, ekonomi, dan sosial di Indonesia. Konsep bela negara mengalami perubahan, mulai dari fokus pada perjuangan fisik melawan penjajah, kemudian bergeser menjadi upaya menjaga keutuhan NKRI dari berbagai ancaman, baik internal maupun eksternal.

Dari konteks tersebut anda diharapkan untuk dapat mengembangkan kalimat itu menjadi beberapa paragraph, yang dimana outputnya dalam format JSON seperti:

{
     "expanded": "Aspek historis bela negara merupakan salah satu dimensi penting dalam memahami konsep bela negara secara menyeluruh. Pemahaman ini tidak hanya sebatas mengingat peristiwa masa lampau, tetapi juga menelisik bagaimana semangat dan nilai-nilai perjuangan masa lalu diwariskan dan diadaptasikan dalam menghadapi tantangan kekinian. Melalui  penelusuran jejak sejarah, kita dapat menyingkap nilai-nilai luhur yang terkandung dalam bela negara,  merunut  evolusi  konsep  bela negara  di  Indonesia,  dan  menarik  pelajaran berharga untuk  menjalani  kehidupan  berbangsa  dan  bernegara  di  era  modern. \n\nBela negara di Indonesia telah berkembang seiring dengan dinamika sejarah bangsa. Pada masa penjajahan, rakyat Indonesia berjuang melawan penjajah dengan berbagai cara, baik secara fisik maupun nonfisik. Perjuangan fisik ditandai dengan perlawanan bersenjata, seperti Perang Diponegoro, Perang Aceh, dan Perang Padri. Sementara itu, perlawanan nonfisik diwujudkan melalui gerakan diplomasi, penyebaran informasi, dan pembangkitan kesadaran nasional. Perjuangan rakyat Indonesia pada masa penjajahan membuktikan bahwa semangat bela negara telah tertanam kuat dalam jiwa bangsa. \n\nSemangat juang rakyat Indonesia pada masa penjajahan  menginspirasi  lahirnya  Sumpah  Pemuda  pada  tahun  1928,  yang  menyatakan tekad  para  pemuda  Indonesia  untuk  bersatu  dan  merdeka.  Sumpah Pemuda menjadi tonggak penting dalam  perjuangan  kemerdekaan,  menyatukan tekad bangsa untuk  meraih  kemerdekaan.  Setelah  Indonesia  merdeka  pada  tahun  1945,  konsep  bela  negara  terus berkembang  seiring  dengan  dinamika  politik,  ekonomi,  dan  sosial  di  Indonesia.\n\nKonsep bela negara setelah kemerdekaan mengalami  perubahan  signifikan.  Pada  masa  perjuangan  kemerdekaan,  fokus  bela negara  terpusat  pada  perjuangan  fisik  melawan  penjajah.  Namun,  setelah Indonesia  merdeka,  konsep  bela  negara  bergeser  menjadi  upaya  menjaga  keutuhan NKRI  dari  berbagai  ancaman,  baik  internal  maupun  eksternal.  Ancaman  internal  yang  dihadapi  Indonesia  antara  lain  konflik  horizontal,  terorisme,  dan  separatisme.  Sementara  itu,  ancaman  eksternal  yang  dihadapi  Indonesia  antara  lain  ancaman  militer,  ekonomi,  dan  ideologi.  Konsep  bela  negara  di masa  kini  memandang  pertahanan  negara  sebagai  suatu  sistem  yang  menyeluruh,  melibatkan  semua  elemen  bangsa,  tidak  hanya  TNI  dan  Polri.\n\nSeiring  dengan  perkembangan  teknologi  dan  informasi,  konsep  bela negara  di  Indonesia  mengalami  adaptasi  baru.  Bela  negara  tidak  hanya  dilakukan  dengan  cara  konvensional,  seperti  perjuangan  fisik,  tetapi  juga  melalui  perjuangan  siber  dan  perjuangan  di  bidang  informasi.  Contohnya  adalah  upaya  menangkal  hoaks  dan  propaganda  negatif  yang  mengancam  keutuhan  NKRI.  Perkembangan  teknologi  informasi  dan  komunikasi  telah  memberikan  peluang  baru  bagi  warga  negara  untuk  berperan  aktif  dalam  bela  negara.  \n\nMelalui  pemahaman  aspek  historis  bela  negara,  kita  dapat  menarik beberapa  pelajaran  berharga.  Pertama,  semangat  bela  negara  merupakan  warisan  luhur  bangsa  Indonesia  yang  harus  terus  dijaga  dan  diwariskan  kepada generasi  berikutnya.  Kedua,  konsep  bela  negara  terus  berkembang  seiring dengan  perkembangan  zaman  dan  tantangan  yang  dihadapi  bangsa.  Ketiga,  bela negara  bukan  hanya  tanggung  jawab  TNI  dan  Polri,  tetapi  juga  tanggung jawab  semua  warga  negara.  Setiap  warga  negara  memiliki  peran  penting  dalam menjaga  keutuhan  NKRI  sesuai  dengan  potensi  dan  keahlian  masing-masing.  \n"
}

Hal yang perlu diperhartikan
- Jangan gunakan formating markdown
                        

Sekarang giliran anda:

{{.task}}
`
