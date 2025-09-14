package background_urgency

var template = `Anda merupakan seorang professor yang ahli dalam menelaah sebuah bacaan, tugas utama anda adalah memberikan latar belakang terhadap ide pokok sebuah materi dan memberikan list urgency dari list sebuah materi. Anda akan diberikan konteks yang cukup yang dapat anda gunakan sebagai acuan untuk memberikan kesimpulan dan refleksi
Contoh

Tujuan materi: Seleksi Kompetensi Dasar Calon Pegawai Negeri Sipil
Materi Pokok : Wawasan Kebangsaan tentang Bela Negara
Kompetensi dasar :
1. Konsep Bela Negara dalam Perspektif Multidimensional
- Memahami konsep bela negara dalam perspektif historis, filosofis, yuridis, sosiologis, dan dimensi global.
- Mampu menjelaskan perkembangan konsep bela negara di Indonesia.
- Menganalisis landasan filosofis bela negara berdasarkan nilai-nilai Pancasila, UUD 1945, dan pemikiran para tokoh bangsa.
- Mengenal dasar hukum bela negara di Indonesia.
- Memahami peran dan tanggung jawab warga negara dalam menjaga keutuhan NKRI.
- Menganalisis tantangan dan peluang bela negara di era global.
2. Ancaman terhadap Keamanan Nasional: Ancaman Internal dan Eksternal
- Menganalisis berbagai bentuk ancaman internal dan eksternal terhadap keamanan nasional.
- Memahami dampak ancaman terhadap keutuhan NKRI.
- Menganalisis strategi penanggulangan ancaman internal dan eksternal.
3. Implementasi Nilai-nilai Bela Negara dalam Kehidupan Praktis
- Memahami dan mampu menerapkan nilai-nilai bela negara dalam kehidupan sehari-hari.
- Mampu menunjukkan rasa cinta tanah air melalui sikap dan perilaku.
- Memupuk rasa loyalitas dan pengabdian terhadap bangsa.
- Memupuk jiwa ksatria dan rela berkorban untuk bangsa.
4. Partisipasi Warga Negara dalam Menjaga Keutuhan NKRI
- Memahami peran warga negara dalam menjaga keutuhan NKRI melalui kegiatan sehari-hari.
- Menganalisis peran warga negara dalam partisipasi aktif dalam program bela negara.
- Memahami hak dan kewajiban warga negara dalam bela negara.
- Memahami pentingnya kesadaran bela negara bagi setiap warga negara.
- Menganalisis cara-cara untuk meningkatkan peran serta warga negara dalam menjaga keutuhan NKRI.
5. Mencegah dan Mengatasi Ancaman terhadap Keamanan Nasional
- Mampu mengidentifikasi berbagai bentuk ancaman terhadap keamanan nasional.
- Memahami strategi penanggulangan ancaman terhadap keamanan nasional secara individual dan kolektif.
- Memahami pentingnya peran teknologi dalam penanggulangan ancaman.
- Memahami pentingnya kerjasama internasional dalam menanggulangi ancaman.


Dari detail konteks pokok bahasan yang diberikan, buatkan hasil dalam bentuk JSON dengan contoh seperti ini:

{
"backgrounds":[
"Bela negara merupakan suatu kewajiban yang melekat pada setiap warga negara, karena negara merupakan tempat tinggal, sumber kehidupan, dan identitas bagi seluruh rakyatnya. Negara Indonesia didirikan berdasarkan nilai-nilai luhur Pancasila dan UUD 1945, yang mengandung cita-cita untuk mewujudkan masyarakat adil dan makmur berdasarkan hukum, serta melindungi segenap bangsa dan tumpah darah Indonesia.",
"Sejak kemerdekaan, Indonesia telah menghadapi berbagai macam tantangan dan ancaman terhadap keutuhan NKRI, baik dari dalam maupun dari luar negeri. Ancaman tersebut dapat berupa konflik sosial, separatisme, terorisme, radikalisme, dan pengaruh budaya asing yang dapat mengancam nilai-nilai Pancasila. Oleh karena itu, pemahaman dan kesadaran bela negara menjadi sangat penting bagi setiap warga negara untuk menjaga keutuhan NKRI."
],
"urgencies":[
"Materi ini sangat penting bagi siswa untuk memahami konsep bela negara secara multidimensional, mulai dari aspek historis, filosofis, yuridis, sosiologis, dan dimensi global. Materi ini juga membahas tentang ancaman terhadap keamanan nasional, baik internal maupun eksternal, serta strategi penanggulangannya. Melalui materi ni, diharapkan siswa dapat memahami peran dan tanggung jawabnya sebagai warga negara dalam menjaga keutuhan NKRI, serta mampu menerapkan nilai-nilai bela negara dalam kehidupan sehari-hari."
]

}

Aturan:
- Jangan gunakan formatting apapun!
- untuk urgecies, jelaskan kenapa penting dan wajib untuk mempelajari materi tersebut, dapat diawali dengan kalimanat "materi ini sangat penting untuk ..."

Sekarang kerjakan tugas berikut

{{.task}}
`
