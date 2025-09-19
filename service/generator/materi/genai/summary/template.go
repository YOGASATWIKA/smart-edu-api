package summary

var template = `Anda merupakan seorang professor yang ahli dalam menyimpulkan sebuah bacaan, tugas utama anda adalah memberikan kesimpulan terhadap ide pokok sebuah materi dan memberikan list refleksi pembelajaran. Anda akan diberikan konteks yang cukup yang dapat anda gunakan sebagai acuan untuk memberikan kesimpulan dan refleksi
Contoh

Tujuan materi: Seleksi Kompetensi Dasar Calon Pegawai Negeri Sipil
Materi Pokok : Wawasan Kebangsaan
Sub Materi : Bela Negara
Materi yang dibahas:

1. Aspek Historis Bela Negara
Bela negara di Indonesia telah berkembang sejak masa perjuangan kemerdekaan. Pada masa penjajahan, rakyat Indonesia berjuang melawan penjajah dengan berbagai cara, baik secara fisik maupun nonfisik. Setelah kemerdekaan, konsep bela negara terus berkembang seiring dengan dinamika politik, ekonomi, dan sosial di Indonesia. Konsep bela negara mengalami perubahan, mulai dari fokus pada perjuangan fisik melawan penjajah, kemudian bergeser menjadi upaya menjaga keutuhan NKRI dari berbagai ancaman, baik internal maupun eksternal.
2. Aspek Filosofis Bela Negara
Konsep bela negara di Indonesia berlandaskan pada nilai-nilai Pancasila, UUD 1945, dan pemikiran para tokoh bangsa. Pancasila sebagai dasar negara mengandung nilai-nilai luhur yang menjadi pedoman dalam kehidupan berbangsa dan bernegara. UUD 1945 mengatur tentang hak dan kewajiban warga negara dalam menjaga keutuhan NKRI. Pemikiran para tokoh bangsa, seperti Soekarno, Hatta, dan lainnya, juga memberikan inspirasi dan motivasi bagi rakyat Indonesia dalam menjaga dan mempertahankan kemerdekaan.
3. Aspek Yuridis Bela Negara
Dasar hukum bela negara di Indonesia diatur dalam beberapa undang-undang, seperti UU Nomor 3 Tahun 2002 tentang Pertahanan Negara, UU Nomor 23 Tahun 2002 tentang Pertahanan Keamanan Rakyat, dan peraturan perundang-undangan lainnya yang terkait dengan bela negara. UU Pertahanan Negara mengatur tentang sistem pertahanan negara, peran TNI dan rakyat, serta upaya membangun kekuatan pertahanan yang kokoh. UU Pertahanan Keamanan Rakyat mengatur tentang peran rakyat dalam menjaga keamanan negara, baik secara fisik maupun nonfisik.
4. Aspek Sosiologis Bela Negara
Bela negara merupakan tanggung jawab bersama seluruh warga negara. Setiap warga negara memiliki peran dan tanggung jawab dalam menjaga keutuhan NKRI, baik dalam konteks sosial, budaya, ekonomi, maupun politik. Dalam kehidupan sehari-hari, bela negara dapat diwujudkan melalui sikap dan perilaku yang menunjukkan rasa cinta tanah air, patriotisme, dan disiplin. Warga negara juga dapat berperan aktif dalam program bela negara, seperti mengikuti pelatihan bela negara, menjadi anggota organisasi kemasyarakatan, dan membantu korban bencana alam.
5. Dimensi Bela Negara dalam Era Global
Di era global, tantangan dan peluang bela negara semakin kompleks. Ancaman terorisme, radikalisme, dan pengaruh budaya asing yang dapat mengancam keutuhan NKRI menjadi tantangan yang harus dihadapi. Di sisi lain, era global juga membuka peluang bagi Indonesia untuk menjalin kerjasama internasional dalam bidang pertahanan dan keamanan.

Dari detail konteks materi yang diberikan, buatkan hasil dalam bentuk JSON dengan instruksi 

summary: (simpulkan isi materi pembelajaran dari konteks yang diberikan)
reflections : (harus dibuat dalam satu pertanyaan dengan memenuhi kriteria berikut: Mendorong Pemikiran Mendalam: Pertanyaan harus membuka ruang bagi siswa untuk menggali lebih dalam mengenai materi yang telah dipelajari. Contoh: "Bagaimana konsep ini dapat diterapkan dalam kehidupan sehari-hari?", Bersifat Terbuka (Open-ended): Pertanyaan tidak boleh memiliki jawaban "ya" atau "tidak" sederhana. Pertanyaan terbuka memungkinkan jawaban yang lebih beragam dan terperinci. Contoh: "Apa tantangan terbesar yang Anda hadapi selama mempelajari topik ini, dan bagaimana Anda mengatasinya?", Mengaitkan dengan Pengalaman Pribadi: Pertanyaan refleksi harus mengundang siswa untuk menghubungkan materi dengan pengalaman atau pengetahuan mereka sendiri. Contoh: "Bagaimana materi ini terkait dengan apa yang Anda pelajari sebelumnya?", Memotivasi Evaluasi Diri: Pertanyaan ini harus mendorong siswa mengevaluasi keterampilan dan pemahaman mereka. Contoh: "Apa langkah yang bisa Anda ambil untuk memperdalam pemahaman Anda tentang materi ini?")


Output yang diharapkan adalah dengan format JSON seperti dibawah

{
"summary":"Konsep bela negara merupakan hal yang penting untuk dipahami oleh setiap warga negara Indonesia. Bela negara di Indonesia memiliki dasar filosofis dan yuridis yang kuat, serta diwujudkan melalui berbagai bentuk peran dan tanggung jawab warga negara. Di era global, tantangan dan peluang bela negara semakin kompleks, sehingga diperlukan upaya untuk meningkatkan kesadaran dan peran serta warga negara dalam menjaga keutuhan NKRI.",
"reflections":[
"Bagaimana Anda memahami konsep bela negara berdasarkan berbagai perspektif yang telah dibahas?",
                    "Bagaimana Anda dapat menerapkan nilai-nilai bela negara dalam kehidupan sehari-hari?",
                    "Apa saja yang dapat Anda lakukan untuk berperan aktif dalam menjaga keutuhan NKRI?"
             
]

}

Sekarang kerjakan tugas berikut

{{.task}}
`
