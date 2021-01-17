# SANCertificate
Create Let's Encrypt signed SAN Certificate using AUTOCERT in GO


Go ile web sunucusu üzerinden kodlama yapanlar için Let's Encrypt tarafından imzalanacak sertifikalar için kütüphaneler kullanılmaktadır. Bunlar arasında benim baktığım kadarıyla HTTP-01 ile sertifika imzalanırken SAN özelliği eklenen bulunmuyor. Benim kullandığım AUTOCERT kütüphanesi üzerinde kaba değişiklik yaparak sertifikalarınıza SAN bilgisini de ekleyebilirsiniz.
